// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: assets/v1/assets.proto

package assetsv1

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

// Validate checks the field values on FetchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FetchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FetchRequestMultiError, or
// nil if none found.
func (m *FetchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FetchRequestMultiError(errors)
	}
	return nil
}

// FetchRequestMultiError is an error wrapping multiple validation errors
// returned by FetchRequest.ValidateAll() if the designated constraints aren't met.
type FetchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchRequestMultiError) AllErrors() []error { return m }

// FetchRequestValidationError is the validation error returned by
// FetchRequest.Validate if the designated constraints aren't met.
type FetchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchRequestValidationError) ErrorName() string { return "FetchRequestValidationError" }

// Error satisfies the builtin error interface
func (e FetchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchRequestValidationError{}

// Validate checks the field values on FetchResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FetchResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FetchResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FetchResponseMultiError, or
// nil if none found.
func (m *FetchResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *FetchResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FetchResponseMultiError(errors)
	}
	return nil
}

// FetchResponseMultiError is an error wrapping multiple validation errors
// returned by FetchResponse.ValidateAll() if the designated constraints
// aren't met.
type FetchResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FetchResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FetchResponseMultiError) AllErrors() []error { return m }

// FetchResponseValidationError is the validation error returned by
// FetchResponse.Validate if the designated constraints aren't met.
type FetchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FetchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FetchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FetchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FetchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FetchResponseValidationError) ErrorName() string { return "FetchResponseValidationError" }

// Error satisfies the builtin error interface
func (e FetchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFetchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FetchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FetchResponseValidationError{}
