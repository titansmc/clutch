package gateway

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"gopkg.in/yaml.v3"

	gatewayv1 "github.com/lyft/clutch/backend/api/config/gateway/v1"
	"github.com/lyft/clutch/backend/middleware/timeouts"
)

type envFiles []string

func (f *envFiles) String() string {
	return strings.Join(*f, ",")
}

func (f *envFiles) Set(value string) error {
	*f = append(*f, value)
	return nil
}

type Flags struct {
	ConfigPath string
	Template   bool
	Validate   bool
	Debug      bool
	EnvFiles   envFiles
}

// Link register the struct vars globally for parsing by the flag library.
func (f *Flags) Link() {
	flag.StringVar(&f.ConfigPath, "c", "clutch-config.yaml", "path to YAML configuration")
	flag.BoolVar(&f.Template, "template", false, "executes go templates on the configuration file")
	flag.BoolVar(&f.Validate, "validate", false, "validates the configuration file and exits")
	flag.BoolVar(&f.Debug, "debug", false, "print the final composed configuration file to stdout")
	flag.Var(&f.EnvFiles, "env", "path to additional .env files to load")
}

// Parse command line arguments.
func ParseFlags() *Flags {
	f := &Flags{}
	f.Link()
	flag.Parse()
	return f
}

func MustReadOrValidateConfig(f *Flags) *gatewayv1.Config {
	// Use a temporary logger to parse the configuration and output.
	tmpLogger := newTmpLogger().With(zap.String("file", f.ConfigPath))

	var cfg gatewayv1.Config
	var seenCfgs []string
	consolidateConfigs(filepath.Dir(f.ConfigPath), filepath.Base(f.ConfigPath), &cfg, f, &seenCfgs)
	if err := cfg.Validate(); err != nil {
		tmpLogger.Fatal("configuration proto validation failed", zap.Error(err))
	}

	if f.Validate {
		if err := ensureUnique(&cfg); err != nil {
			tmpLogger.Fatal("configuration validation failed", zap.Error(err))
		}

		tmpLogger.Info("configuration validation was successful")
		os.Exit(0)
	}

	if f.Debug {
		json, err := protojson.Marshal(&cfg)
		if err != nil {
			tmpLogger.Fatal("failed to cast configuration file to json", zap.Error(err))
		}
		fmt.Println(string(json))
		os.Exit(0)
	}

	return &cfg
}

// Ensure gateway config has unique modules and service entries
func ensureUnique(cfg *gatewayv1.Config) error {
	svcs := make(map[string]bool)
	for _, s := range cfg.GetServices() {
		if _, seen := svcs[s.GetName()]; seen {
			return fmt.Errorf("duplicate service found: %s", s.GetName())
		}
		svcs[s.GetName()] = true
	}

	mods := make(map[string]bool)
	for _, m := range cfg.GetModules() {
		if _, seen := mods[m.GetName()]; seen {
			return fmt.Errorf("duplicate module found: %s", m.GetName())
		}
		mods[m.GetName()] = true
	}

	rlvs := make(map[string]bool)
	for _, r := range cfg.GetResolvers() {
		if _, seen := rlvs[r.GetName()]; seen {
			return fmt.Errorf("duplicate resolver found: %s", r.GetName())
		}
		rlvs[r.GetName()] = true
	}
	return nil
}

func contains(s *[]string, str string) bool {
	for _, v := range *s {
		if v == str {
			return true
		}
	}

	return false
}

func consolidateConfigs(cfgBaseDir, cfgFile string, cfg *gatewayv1.Config, f *Flags, seen *[]string) {
	// Use a temporary logger to parse the configuration and output.
	cfgPath := filepath.Join(cfgBaseDir, cfgFile)
	tmpLogger := newTmpLogger().With(zap.String("file", cfgPath))

	if contains(seen, cfgPath) {
		tmpLogger.Warn("ignoring duplicate extended config")
		return
	}

	var curCfg gatewayv1.Config
	if err := parseFile(cfgPath, &curCfg, f.Template); err != nil {
		tmpLogger.Fatal("parsing configuration failed", zap.Error(err))
	}

	*seen = append(*seen, cfgPath)
	if len(curCfg.Extends) == 0 {
		proto.Merge(cfg, &curCfg)
		return
	}

	for _, c := range curCfg.Extends {
		if c == cfgPath {
			continue
		}
		consolidateConfigs(filepath.Dir(cfgPath), c, cfg, f, seen)
	}
	proto.Merge(cfg, &curCfg)
}

func executeTemplate(contents []byte) ([]byte, error) {
	tmpl := template.New("config").Funcs(map[string]interface{}{
		"getenv": os.Getenv,
		"getboolenv": func(key string) bool {
			b, _ := strconv.ParseBool(os.Getenv(key))
			return b
		},
	})

	tmpl, err := tmpl.Parse(string(contents))
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, nil); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func parseFile(path string, pb proto.Message, template bool) error {
	// Get absolute path representation for better error message in case file not found.
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// Read file.
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Execute templates if enabled.
	if template {
		contents, err = executeTemplate(contents)
		if err != nil {
			return err
		}
	}

	/*
		We want to support defining Go templates in the clutch-config that can be executed at
		a later time.
		Two issues that needed to be addressed:
		1) os.ExpandEnv removes the $ and there isn't a way to provide an escape
		(open upstream issue: https://github.com/golang/go/issues/43482)
		2) the config itself can be executed as a template, and so the nested templates would also
		be executed.

		Solution is to use Clutch-specific templating tokens in the config that are then replaced
		with the Go Template syntax
		1) $$ in lieu of $
		2) [[ ]] in lieu of {{ }}
	*/

	tokenContent := bulkReplaceTemplateTokens(string(contents))

	// Interpolate environment variables
	expandedContent := os.ExpandEnv(tokenContent)

	contents = []byte(replaceVarTemplateToken(expandedContent))

	return parseYAML(contents, pb)
}

// swaps the Clutch "Action" tokens for the Go Template "Action" tokens
// swaps the dollar signs with other characters, otherwise os.ExpandEnv
// would remove the dollar signs
func bulkReplaceTemplateTokens(data string) string {
	sanitize := strings.NewReplacer(
		"[[", "{{",
		"]]", "}}",
		"$$", "@#@",
	)
	return sanitize.Replace(data)
}

// swaps the Clutch variable token with the Go Template variable token
func replaceVarTemplateToken(data string) string {
	return strings.ReplaceAll(data, "@#@", "$")
}

func parseYAML(contents []byte, pb proto.Message) error {
	// Decode YAML.
	var rawConfig map[string]interface{}
	if err := yaml.Unmarshal(contents, &rawConfig); err != nil {
		return err
	}

	// Encode YAML to JSON.
	rawJSON, err := json.Marshal(rawConfig)
	if err != nil {
		return err
	}

	// Unmarshal JSON to proto object.
	if err := protojson.Unmarshal(rawJSON, pb); err != nil {
		return err
	}

	// All good!
	return nil
}

func newLogger(msg *gatewayv1.Logger) (*zap.Logger, error) {
	return newLoggerWithCore(msg, nil)
}

func newLoggerWithCore(msg *gatewayv1.Logger, zapCore zapcore.Core) (*zap.Logger, error) {
	var c zap.Config
	var opts []zap.Option
	if msg.GetPretty() {
		c = zap.NewDevelopmentConfig()
		opts = append(opts, zap.AddStacktrace(zap.ErrorLevel))
	} else {
		c = zap.NewProductionConfig()
	}

	level := zap.NewAtomicLevel()

	levelName := "INFO"
	if msg.Level != gatewayv1.Logger_UNSPECIFIED {
		levelName = msg.Level.String()
	}

	if err := level.UnmarshalText([]byte(levelName)); err != nil {
		return nil, fmt.Errorf("could not parse log level %s", msg.Level.String())
	}
	c.Level = level

	logger, err := c.Build(opts...)
	if err != nil {
		return nil, err
	}

	// If zapCore is set, create a new logger, this is currently only used in tests.
	if zapCore != nil {
		logger = zap.New(zapCore, opts...)
	}

	if len(msg.Namespace) > 0 {
		logger = logger.With(zap.Namespace(msg.Namespace))
	}

	return logger, nil
}

func newTmpLogger() *zap.Logger {
	c := zap.NewProductionConfig()
	c.DisableStacktrace = true
	l, err := c.Build()
	if err != nil {
		panic(err)
	}
	return l
}

type validator interface {
	Validate() error
}

// Returns maximum timeout, where 0 is considered maximum (i.e. no timeout).
func computeMaximumTimeout(cfg *gatewayv1.Timeouts) time.Duration {
	if cfg == nil {
		return timeouts.DefaultTimeout
	}

	ret := cfg.Default.AsDuration()
	for _, e := range cfg.Overrides {
		override := e.Timeout.AsDuration()
		if ret == 0 || override == 0 {
			return 0
		}

		if override > ret {
			ret = override
		}
	}

	return ret
}

func validateAny(a *anypb.Any) error {
	if a == nil {
		return nil
	}

	m, err := a.UnmarshalNew()
	if err != nil {
		return err
	}

	if v, ok := m.(validator); ok {
		return v.Validate()
	}
	return nil
}
