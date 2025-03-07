// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoytriage/v1/envoytriage_api.proto

package envoytriagev1

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
)

// Validate checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadRequestMultiError, or
// nil if none found.
func (m *ReadRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetOperations() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReadRequestValidationError{
						field:  fmt.Sprintf("Operations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReadRequestValidationError{
						field:  fmt.Sprintf("Operations[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReadRequestValidationError{
					field:  fmt.Sprintf("Operations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ReadRequestMultiError(errors)
	}
	return nil
}

// ReadRequestMultiError is an error wrapping multiple validation errors
// returned by ReadRequest.ValidateAll() if the designated constraints aren't met.
type ReadRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadRequestMultiError) AllErrors() []error { return m }

// ReadRequestValidationError is the validation error returned by
// ReadRequest.Validate if the designated constraints aren't met.
type ReadRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadRequestValidationError) ErrorName() string { return "ReadRequestValidationError" }

// Error satisfies the builtin error interface
func (e ReadRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadRequestValidationError{}

// Validate checks the field values on ReadOperation with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadOperation) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadOperation with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadOperationMultiError, or
// nil if none found.
func (m *ReadOperation) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadOperation) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetAddress() == nil {
		err := ReadOperationValidationError{
			field:  "Address",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReadOperationValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadOperationValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadOperationValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetInclude()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ReadOperationValidationError{
					field:  "Include",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ReadOperationValidationError{
					field:  "Include",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInclude()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReadOperationValidationError{
				field:  "Include",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ReadOperationMultiError(errors)
	}
	return nil
}

// ReadOperationMultiError is an error wrapping multiple validation errors
// returned by ReadOperation.ValidateAll() if the designated constraints
// aren't met.
type ReadOperationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadOperationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadOperationMultiError) AllErrors() []error { return m }

// ReadOperationValidationError is the validation error returned by
// ReadOperation.Validate if the designated constraints aren't met.
type ReadOperationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadOperationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadOperationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadOperationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadOperationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadOperationValidationError) ErrorName() string { return "ReadOperationValidationError" }

// Error satisfies the builtin error interface
func (e ReadOperationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadOperation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadOperationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadOperationValidationError{}

// Validate checks the field values on ReadResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ReadResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ReadResponseMultiError, or
// nil if none found.
func (m *ReadResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ReadResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ReadResponseValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReadResponseValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ReadResponseMultiError(errors)
	}
	return nil
}

// ReadResponseMultiError is an error wrapping multiple validation errors
// returned by ReadResponse.ValidateAll() if the designated constraints aren't met.
type ReadResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadResponseMultiError) AllErrors() []error { return m }

// ReadResponseValidationError is the validation error returned by
// ReadResponse.Validate if the designated constraints aren't met.
type ReadResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadResponseValidationError) ErrorName() string { return "ReadResponseValidationError" }

// Error satisfies the builtin error interface
func (e ReadResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadResponseValidationError{}

// Validate checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Address) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Address with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in AddressMultiError, or nil if none found.
func (m *Address) ValidateAll() error {
	return m.validate(true)
}

func (m *Address) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateHostname(m.GetHost()); err != nil {
		if ip := net.ParseIP(m.GetHost()); ip == nil {
			err := AddressValidationError{
				field:  "Host",
				reason: "value must be a valid hostname, or ip address",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	if m.GetPort() > 65535 {
		err := AddressValidationError{
			field:  "Port",
			reason: "value must be less than or equal to 65535",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AddressMultiError(errors)
	}
	return nil
}

func (m *Address) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// AddressMultiError is an error wrapping multiple validation errors returned
// by Address.ValidateAll() if the designated constraints aren't met.
type AddressMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddressMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddressMultiError) AllErrors() []error { return m }

// AddressValidationError is the validation error returned by Address.Validate
// if the designated constraints aren't met.
type AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddressValidationError) ErrorName() string { return "AddressValidationError" }

// Error satisfies the builtin error interface
func (e AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddress.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddressValidationError{}

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ResultMultiError, or nil if none found.
func (m *Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResultValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResultValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetNodeMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResultValidationError{
					field:  "NodeMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResultValidationError{
					field:  "NodeMetadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNodeMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultValidationError{
				field:  "NodeMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOutput()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ResultValidationError{
					field:  "Output",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ResultValidationError{
					field:  "Output",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOutput()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultValidationError{
				field:  "Output",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ResultMultiError(errors)
	}
	return nil
}

// ResultMultiError is an error wrapping multiple validation errors returned by
// Result.ValidateAll() if the designated constraints aren't met.
type ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultMultiError) AllErrors() []error { return m }

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}

// Validate checks the field values on NodeMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NodeMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NodeMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NodeMetadataMultiError, or
// nil if none found.
func (m *NodeMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *NodeMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ServiceNode

	// no validation rules for ServiceCluster

	// no validation rules for ServiceZone

	// no validation rules for Version

	if len(errors) > 0 {
		return NodeMetadataMultiError(errors)
	}
	return nil
}

// NodeMetadataMultiError is an error wrapping multiple validation errors
// returned by NodeMetadata.ValidateAll() if the designated constraints aren't met.
type NodeMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NodeMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NodeMetadataMultiError) AllErrors() []error { return m }

// NodeMetadataValidationError is the validation error returned by
// NodeMetadata.Validate if the designated constraints aren't met.
type NodeMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeMetadataValidationError) ErrorName() string { return "NodeMetadataValidationError" }

// Error satisfies the builtin error interface
func (e NodeMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNodeMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeMetadataValidationError{}

// Validate checks the field values on ReadOperation_Include with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ReadOperation_Include) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ReadOperation_Include with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ReadOperation_IncludeMultiError, or nil if none found.
func (m *ReadOperation_Include) ValidateAll() error {
	return m.validate(true)
}

func (m *ReadOperation_Include) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Clusters

	// no validation rules for ConfigDump

	// no validation rules for Listeners

	// no validation rules for Runtime

	// no validation rules for Stats

	// no validation rules for ServerInfo

	if len(errors) > 0 {
		return ReadOperation_IncludeMultiError(errors)
	}
	return nil
}

// ReadOperation_IncludeMultiError is an error wrapping multiple validation
// errors returned by ReadOperation_Include.ValidateAll() if the designated
// constraints aren't met.
type ReadOperation_IncludeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ReadOperation_IncludeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ReadOperation_IncludeMultiError) AllErrors() []error { return m }

// ReadOperation_IncludeValidationError is the validation error returned by
// ReadOperation_Include.Validate if the designated constraints aren't met.
type ReadOperation_IncludeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReadOperation_IncludeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReadOperation_IncludeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReadOperation_IncludeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReadOperation_IncludeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReadOperation_IncludeValidationError) ErrorName() string {
	return "ReadOperation_IncludeValidationError"
}

// Error satisfies the builtin error interface
func (e ReadOperation_IncludeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReadOperation_Include.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReadOperation_IncludeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReadOperation_IncludeValidationError{}

// Validate checks the field values on Result_Output with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Result_Output) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Result_Output with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Result_OutputMultiError, or
// nil if none found.
func (m *Result_Output) ValidateAll() error {
	return m.validate(true)
}

func (m *Result_Output) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetClusters()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Clusters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Clusters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClusters()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Result_OutputValidationError{
				field:  "Clusters",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetConfigDump()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "ConfigDump",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "ConfigDump",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfigDump()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Result_OutputValidationError{
				field:  "ConfigDump",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetListeners()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Listeners",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Listeners",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetListeners()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Result_OutputValidationError{
				field:  "Listeners",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRuntime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Runtime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Runtime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Result_OutputValidationError{
				field:  "Runtime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Result_OutputValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetServerInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "ServerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Result_OutputValidationError{
					field:  "ServerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetServerInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Result_OutputValidationError{
				field:  "ServerInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Result_OutputMultiError(errors)
	}
	return nil
}

// Result_OutputMultiError is an error wrapping multiple validation errors
// returned by Result_Output.ValidateAll() if the designated constraints
// aren't met.
type Result_OutputMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Result_OutputMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Result_OutputMultiError) AllErrors() []error { return m }

// Result_OutputValidationError is the validation error returned by
// Result_Output.Validate if the designated constraints aren't met.
type Result_OutputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Result_OutputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Result_OutputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Result_OutputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Result_OutputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Result_OutputValidationError) ErrorName() string { return "Result_OutputValidationError" }

// Error satisfies the builtin error interface
func (e Result_OutputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult_Output.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Result_OutputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Result_OutputValidationError{}
