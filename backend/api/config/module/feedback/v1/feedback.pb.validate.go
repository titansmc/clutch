// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/module/feedback/v1/feedback.proto

package feedbackmodv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	feedbackv1 "github.com/lyft/clutch/backend/api/feedback/v1"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = feedbackv1.Origin(0)
)

// Validate checks the field values on Survey with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Survey) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Survey with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SurveyMultiError, or nil if none found.
func (m *Survey) ValidateAll() error {
	return m.validate(true)
}

func (m *Survey) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetPrompt()) < 1 {
		err := SurveyValidationError{
			field:  "Prompt",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for FreeformPrompt

	if len(m.GetRatingLabels()) < 2 {
		err := SurveyValidationError{
			field:  "RatingLabels",
			reason: "value must contain at least 2 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetRatingLabels() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SurveyValidationError{
						field:  fmt.Sprintf("RatingLabels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SurveyValidationError{
						field:  fmt.Sprintf("RatingLabels[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SurveyValidationError{
					field:  fmt.Sprintf("RatingLabels[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SurveyMultiError(errors)
	}
	return nil
}

// SurveyMultiError is an error wrapping multiple validation errors returned by
// Survey.ValidateAll() if the designated constraints aren't met.
type SurveyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SurveyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SurveyMultiError) AllErrors() []error { return m }

// SurveyValidationError is the validation error returned by Survey.Validate if
// the designated constraints aren't met.
type SurveyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SurveyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SurveyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SurveyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SurveyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SurveyValidationError) ErrorName() string { return "SurveyValidationError" }

// Error satisfies the builtin error interface
func (e SurveyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSurvey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SurveyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SurveyValidationError{}

// Validate checks the field values on SurveyOrigin with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SurveyOrigin) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SurveyOrigin with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SurveyOriginMultiError, or
// nil if none found.
func (m *SurveyOrigin) ValidateAll() error {
	return m.validate(true)
}

func (m *SurveyOrigin) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _SurveyOrigin_Origin_NotInLookup[m.GetOrigin()]; ok {
		err := SurveyOriginValidationError{
			field:  "Origin",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := feedbackv1.Origin_name[int32(m.GetOrigin())]; !ok {
		err := SurveyOriginValidationError{
			field:  "Origin",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSurvey() == nil {
		err := SurveyOriginValidationError{
			field:  "Survey",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSurvey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SurveyOriginValidationError{
					field:  "Survey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SurveyOriginValidationError{
					field:  "Survey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSurvey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SurveyOriginValidationError{
				field:  "Survey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SurveyOriginMultiError(errors)
	}
	return nil
}

// SurveyOriginMultiError is an error wrapping multiple validation errors
// returned by SurveyOrigin.ValidateAll() if the designated constraints aren't met.
type SurveyOriginMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SurveyOriginMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SurveyOriginMultiError) AllErrors() []error { return m }

// SurveyOriginValidationError is the validation error returned by
// SurveyOrigin.Validate if the designated constraints aren't met.
type SurveyOriginValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SurveyOriginValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SurveyOriginValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SurveyOriginValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SurveyOriginValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SurveyOriginValidationError) ErrorName() string { return "SurveyOriginValidationError" }

// Error satisfies the builtin error interface
func (e SurveyOriginValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSurveyOrigin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SurveyOriginValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SurveyOriginValidationError{}

var _SurveyOrigin_Origin_NotInLookup = map[feedbackv1.Origin]struct{}{
	0: {},
}

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfigMultiError, or nil if none found.
func (m *Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetOrigins()) < 1 {
		err := ConfigValidationError{
			field:  "Origins",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetOrigins() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("Origins[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConfigValidationError{
						field:  fmt.Sprintf("Origins[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("Origins[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}
	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.ValidateAll() if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}
