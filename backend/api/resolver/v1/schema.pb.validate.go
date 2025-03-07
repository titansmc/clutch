// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: resolver/v1/schema.proto

package resolverv1

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

// Validate checks the field values on StringField with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StringField) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StringField with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StringFieldMultiError, or
// nil if none found.
func (m *StringField) ValidateAll() error {
	return m.validate(true)
}

func (m *StringField) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Placeholder

	// no validation rules for DefaultValue

	if len(errors) > 0 {
		return StringFieldMultiError(errors)
	}
	return nil
}

// StringFieldMultiError is an error wrapping multiple validation errors
// returned by StringField.ValidateAll() if the designated constraints aren't met.
type StringFieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StringFieldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StringFieldMultiError) AllErrors() []error { return m }

// StringFieldValidationError is the validation error returned by
// StringField.Validate if the designated constraints aren't met.
type StringFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringFieldValidationError) ErrorName() string { return "StringFieldValidationError" }

// Error satisfies the builtin error interface
func (e StringFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringFieldValidationError{}

// Validate checks the field values on Option with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Option) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Option with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OptionMultiError, or nil if none found.
func (m *Option) ValidateAll() error {
	return m.validate(true)
}

func (m *Option) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisplayName

	switch m.Value.(type) {

	case *Option_StringValue:
		// no validation rules for StringValue

	}

	if len(errors) > 0 {
		return OptionMultiError(errors)
	}
	return nil
}

// OptionMultiError is an error wrapping multiple validation errors returned by
// Option.ValidateAll() if the designated constraints aren't met.
type OptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionMultiError) AllErrors() []error { return m }

// OptionValidationError is the validation error returned by Option.Validate if
// the designated constraints aren't met.
type OptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionValidationError) ErrorName() string { return "OptionValidationError" }

// Error satisfies the builtin error interface
func (e OptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionValidationError{}

// Validate checks the field values on OptionField with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OptionField) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OptionField with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OptionFieldMultiError, or
// nil if none found.
func (m *OptionField) ValidateAll() error {
	return m.validate(true)
}

func (m *OptionField) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IncludeAllOption

	for idx, item := range m.GetOptions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OptionFieldValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OptionFieldValidationError{
						field:  fmt.Sprintf("Options[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OptionFieldValidationError{
					field:  fmt.Sprintf("Options[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OptionFieldMultiError(errors)
	}
	return nil
}

// OptionFieldMultiError is an error wrapping multiple validation errors
// returned by OptionField.ValidateAll() if the designated constraints aren't met.
type OptionFieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OptionFieldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OptionFieldMultiError) AllErrors() []error { return m }

// OptionFieldValidationError is the validation error returned by
// OptionField.Validate if the designated constraints aren't met.
type OptionFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OptionFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OptionFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OptionFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OptionFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OptionFieldValidationError) ErrorName() string { return "OptionFieldValidationError" }

// Error satisfies the builtin error interface
func (e OptionFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOptionField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OptionFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OptionFieldValidationError{}

// Validate checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Field) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FieldMultiError, or nil if none found.
func (m *Field) ValidateAll() error {
	return m.validate(true)
}

func (m *Field) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FieldValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FieldValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FieldValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FieldMultiError(errors)
	}
	return nil
}

// FieldMultiError is an error wrapping multiple validation errors returned by
// Field.ValidateAll() if the designated constraints aren't met.
type FieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldMultiError) AllErrors() []error { return m }

// FieldValidationError is the validation error returned by Field.Validate if
// the designated constraints aren't met.
type FieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldValidationError) ErrorName() string { return "FieldValidationError" }

// Error satisfies the builtin error interface
func (e FieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldValidationError{}

// Validate checks the field values on FieldMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FieldMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FieldMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FieldMetadataMultiError, or
// nil if none found.
func (m *FieldMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *FieldMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisplayName

	// no validation rules for Required

	switch m.Type.(type) {

	case *FieldMetadata_StringField:

		if all {
			switch v := interface{}(m.GetStringField()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FieldMetadataValidationError{
						field:  "StringField",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FieldMetadataValidationError{
						field:  "StringField",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetStringField()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldMetadataValidationError{
					field:  "StringField",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *FieldMetadata_OptionField:

		if all {
			switch v := interface{}(m.GetOptionField()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FieldMetadataValidationError{
						field:  "OptionField",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FieldMetadataValidationError{
						field:  "OptionField",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOptionField()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FieldMetadataValidationError{
					field:  "OptionField",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return FieldMetadataMultiError(errors)
	}
	return nil
}

// FieldMetadataMultiError is an error wrapping multiple validation errors
// returned by FieldMetadata.ValidateAll() if the designated constraints
// aren't met.
type FieldMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldMetadataMultiError) AllErrors() []error { return m }

// FieldMetadataValidationError is the validation error returned by
// FieldMetadata.Validate if the designated constraints aren't met.
type FieldMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldMetadataValidationError) ErrorName() string { return "FieldMetadataValidationError" }

// Error satisfies the builtin error interface
func (e FieldMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFieldMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldMetadataValidationError{}

// Validate checks the field values on SearchMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SearchMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SearchMetadataMultiError,
// or nil if none found.
func (m *SearchMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enabled

	// no validation rules for AutocompleteEnabled

	if len(errors) > 0 {
		return SearchMetadataMultiError(errors)
	}
	return nil
}

// SearchMetadataMultiError is an error wrapping multiple validation errors
// returned by SearchMetadata.ValidateAll() if the designated constraints
// aren't met.
type SearchMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchMetadataMultiError) AllErrors() []error { return m }

// SearchMetadataValidationError is the validation error returned by
// SearchMetadata.Validate if the designated constraints aren't met.
type SearchMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchMetadataValidationError) ErrorName() string { return "SearchMetadataValidationError" }

// Error satisfies the builtin error interface
func (e SearchMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchMetadataValidationError{}

// Validate checks the field values on SchemaMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SchemaMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SchemaMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SchemaMetadataMultiError,
// or nil if none found.
func (m *SchemaMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *SchemaMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisplayName

	// no validation rules for Searchable

	if all {
		switch v := interface{}(m.GetSearch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SchemaMetadataValidationError{
					field:  "Search",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SchemaMetadataValidationError{
					field:  "Search",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSearch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaMetadataValidationError{
				field:  "Search",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SchemaMetadataMultiError(errors)
	}
	return nil
}

// SchemaMetadataMultiError is an error wrapping multiple validation errors
// returned by SchemaMetadata.ValidateAll() if the designated constraints
// aren't met.
type SchemaMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SchemaMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SchemaMetadataMultiError) AllErrors() []error { return m }

// SchemaMetadataValidationError is the validation error returned by
// SchemaMetadata.Validate if the designated constraints aren't met.
type SchemaMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaMetadataValidationError) ErrorName() string { return "SchemaMetadataValidationError" }

// Error satisfies the builtin error interface
func (e SchemaMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchemaMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaMetadataValidationError{}

// Validate checks the field values on Schema with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Schema) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Schema with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SchemaMultiError, or nil if none found.
func (m *Schema) ValidateAll() error {
	return m.validate(true)
}

func (m *Schema) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TypeUrl

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SchemaValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SchemaValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFields() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SchemaValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SchemaValidationError{
						field:  fmt.Sprintf("Fields[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SchemaValidationError{
					field:  fmt.Sprintf("Fields[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetError()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SchemaValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SchemaValidationError{
					field:  "Error",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SchemaValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SchemaMultiError(errors)
	}
	return nil
}

// SchemaMultiError is an error wrapping multiple validation errors returned by
// Schema.ValidateAll() if the designated constraints aren't met.
type SchemaMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SchemaMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SchemaMultiError) AllErrors() []error { return m }

// SchemaValidationError is the validation error returned by Schema.Validate if
// the designated constraints aren't met.
type SchemaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SchemaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SchemaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SchemaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SchemaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SchemaValidationError) ErrorName() string { return "SchemaValidationError" }

// Error satisfies the builtin error interface
func (e SchemaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSchema.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SchemaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SchemaValidationError{}
