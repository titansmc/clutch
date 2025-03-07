// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: feedback/v1/feedback.proto

package feedbackv1

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

// Validate checks the field values on RatingLabel with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RatingLabel) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RatingLabel with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RatingLabelMultiError, or
// nil if none found.
func (m *RatingLabel) ValidateAll() error {
	return m.validate(true)
}

func (m *RatingLabel) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetLabel()) < 1 {
		err := RatingLabelValidationError{
			field:  "Label",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	switch m.Type.(type) {

	case *RatingLabel_Emoji:

		if _, ok := _RatingLabel_Emoji_NotInLookup[m.GetEmoji()]; ok {
			err := RatingLabelValidationError{
				field:  "Emoji",
				reason: "value must not be in list [0]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if _, ok := EmojiRating_name[int32(m.GetEmoji())]; !ok {
			err := RatingLabelValidationError{
				field:  "Emoji",
				reason: "value must be one of the defined enum values",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		err := RatingLabelValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return RatingLabelMultiError(errors)
	}
	return nil
}

// RatingLabelMultiError is an error wrapping multiple validation errors
// returned by RatingLabel.ValidateAll() if the designated constraints aren't met.
type RatingLabelMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RatingLabelMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RatingLabelMultiError) AllErrors() []error { return m }

// RatingLabelValidationError is the validation error returned by
// RatingLabel.Validate if the designated constraints aren't met.
type RatingLabelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RatingLabelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RatingLabelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RatingLabelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RatingLabelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RatingLabelValidationError) ErrorName() string { return "RatingLabelValidationError" }

// Error satisfies the builtin error interface
func (e RatingLabelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRatingLabel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RatingLabelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RatingLabelValidationError{}

var _RatingLabel_Emoji_NotInLookup = map[EmojiRating]struct{}{
	0: {},
}

// Validate checks the field values on RatingScale with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RatingScale) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RatingScale with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RatingScaleMultiError, or
// nil if none found.
func (m *RatingScale) ValidateAll() error {
	return m.validate(true)
}

func (m *RatingScale) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.Type.(type) {

	case *RatingScale_Emoji:

		if _, ok := _RatingScale_Emoji_NotInLookup[m.GetEmoji()]; ok {
			err := RatingScaleValidationError{
				field:  "Emoji",
				reason: "value must not be in list [0]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if _, ok := EmojiRating_name[int32(m.GetEmoji())]; !ok {
			err := RatingScaleValidationError{
				field:  "Emoji",
				reason: "value must be one of the defined enum values",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	default:
		err := RatingScaleValidationError{
			field:  "Type",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return RatingScaleMultiError(errors)
	}
	return nil
}

// RatingScaleMultiError is an error wrapping multiple validation errors
// returned by RatingScale.ValidateAll() if the designated constraints aren't met.
type RatingScaleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RatingScaleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RatingScaleMultiError) AllErrors() []error { return m }

// RatingScaleValidationError is the validation error returned by
// RatingScale.Validate if the designated constraints aren't met.
type RatingScaleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RatingScaleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RatingScaleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RatingScaleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RatingScaleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RatingScaleValidationError) ErrorName() string { return "RatingScaleValidationError" }

// Error satisfies the builtin error interface
func (e RatingScaleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRatingScale.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RatingScaleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RatingScaleValidationError{}

var _RatingScale_Emoji_NotInLookup = map[EmojiRating]struct{}{
	0: {},
}

// Validate checks the field values on GetSurveysRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetSurveysRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSurveysRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSurveysRequestMultiError, or nil if none found.
func (m *GetSurveysRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSurveysRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetOrigins()) < 1 {
		err := GetSurveysRequestValidationError{
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

		if _, ok := _GetSurveysRequest_Origins_NotInLookup[item]; ok {
			err := GetSurveysRequestValidationError{
				field:  fmt.Sprintf("Origins[%v]", idx),
				reason: "value must not be in list [0]",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if _, ok := Origin_name[int32(item)]; !ok {
			err := GetSurveysRequestValidationError{
				field:  fmt.Sprintf("Origins[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return GetSurveysRequestMultiError(errors)
	}
	return nil
}

// GetSurveysRequestMultiError is an error wrapping multiple validation errors
// returned by GetSurveysRequest.ValidateAll() if the designated constraints
// aren't met.
type GetSurveysRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSurveysRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSurveysRequestMultiError) AllErrors() []error { return m }

// GetSurveysRequestValidationError is the validation error returned by
// GetSurveysRequest.Validate if the designated constraints aren't met.
type GetSurveysRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSurveysRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSurveysRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSurveysRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSurveysRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSurveysRequestValidationError) ErrorName() string {
	return "GetSurveysRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSurveysRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSurveysRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSurveysRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSurveysRequestValidationError{}

var _GetSurveysRequest_Origins_NotInLookup = map[Origin]struct{}{
	0: {},
}

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

	// no validation rules for Prompt

	// no validation rules for FreeformPrompt

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

// Validate checks the field values on GetSurveysResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSurveysResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSurveysResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSurveysResponseMultiError, or nil if none found.
func (m *GetSurveysResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSurveysResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetOriginSurvey()))
		i := 0
		for key := range m.GetOriginSurvey() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetOriginSurvey()[key]
			_ = val

			// no validation rules for OriginSurvey[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, GetSurveysResponseValidationError{
							field:  fmt.Sprintf("OriginSurvey[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, GetSurveysResponseValidationError{
							field:  fmt.Sprintf("OriginSurvey[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return GetSurveysResponseValidationError{
						field:  fmt.Sprintf("OriginSurvey[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return GetSurveysResponseMultiError(errors)
	}
	return nil
}

// GetSurveysResponseMultiError is an error wrapping multiple validation errors
// returned by GetSurveysResponse.ValidateAll() if the designated constraints
// aren't met.
type GetSurveysResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSurveysResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSurveysResponseMultiError) AllErrors() []error { return m }

// GetSurveysResponseValidationError is the validation error returned by
// GetSurveysResponse.Validate if the designated constraints aren't met.
type GetSurveysResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSurveysResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSurveysResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSurveysResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSurveysResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSurveysResponseValidationError) ErrorName() string {
	return "GetSurveysResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSurveysResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSurveysResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSurveysResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSurveysResponseValidationError{}

// Validate checks the field values on FeedbackMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *FeedbackMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FeedbackMetadata with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FeedbackMetadataMultiError, or nil if none found.
func (m *FeedbackMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *FeedbackMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _FeedbackMetadata_Origin_NotInLookup[m.GetOrigin()]; ok {
		err := FeedbackMetadataValidationError{
			field:  "Origin",
			reason: "value must not be in list [0]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Origin_name[int32(m.GetOrigin())]; !ok {
		err := FeedbackMetadataValidationError{
			field:  "Origin",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSurvey() == nil {
		err := FeedbackMetadataValidationError{
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
				errors = append(errors, FeedbackMetadataValidationError{
					field:  "Survey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FeedbackMetadataValidationError{
					field:  "Survey",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSurvey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedbackMetadataValidationError{
				field:  "Survey",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UserSubmitted

	// no validation rules for UrlSearchParams

	if len(errors) > 0 {
		return FeedbackMetadataMultiError(errors)
	}
	return nil
}

// FeedbackMetadataMultiError is an error wrapping multiple validation errors
// returned by FeedbackMetadata.ValidateAll() if the designated constraints
// aren't met.
type FeedbackMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeedbackMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeedbackMetadataMultiError) AllErrors() []error { return m }

// FeedbackMetadataValidationError is the validation error returned by
// FeedbackMetadata.Validate if the designated constraints aren't met.
type FeedbackMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedbackMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedbackMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedbackMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedbackMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedbackMetadataValidationError) ErrorName() string { return "FeedbackMetadataValidationError" }

// Error satisfies the builtin error interface
func (e FeedbackMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedbackMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedbackMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedbackMetadataValidationError{}

var _FeedbackMetadata_Origin_NotInLookup = map[Origin]struct{}{
	0: {},
}

// Validate checks the field values on Feedback with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Feedback) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Feedback with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FeedbackMultiError, or nil
// if none found.
func (m *Feedback) ValidateAll() error {
	return m.validate(true)
}

func (m *Feedback) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetFeedbackType()) < 1 {
		err := FeedbackValidationError{
			field:  "FeedbackType",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetRatingLabel()) < 1 {
		err := FeedbackValidationError{
			field:  "RatingLabel",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRatingScale() == nil {
		err := FeedbackValidationError{
			field:  "RatingScale",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRatingScale()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FeedbackValidationError{
					field:  "RatingScale",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FeedbackValidationError{
					field:  "RatingScale",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRatingScale()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FeedbackValidationError{
				field:  "RatingScale",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FreeformResponse

	if len(errors) > 0 {
		return FeedbackMultiError(errors)
	}
	return nil
}

// FeedbackMultiError is an error wrapping multiple validation errors returned
// by Feedback.ValidateAll() if the designated constraints aren't met.
type FeedbackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeedbackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeedbackMultiError) AllErrors() []error { return m }

// FeedbackValidationError is the validation error returned by
// Feedback.Validate if the designated constraints aren't met.
type FeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedbackValidationError) ErrorName() string { return "FeedbackValidationError" }

// Error satisfies the builtin error interface
func (e FeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedbackValidationError{}

// Validate checks the field values on SubmitFeedbackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubmitFeedbackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitFeedbackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubmitFeedbackRequestMultiError, or nil if none found.
func (m *SubmitFeedbackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitFeedbackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) != 36 {
		err := SubmitFeedbackRequestValidationError{
			field:  "Id",
			reason: "value length must be 36 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(m.GetUserId()) < 1 {
		err := SubmitFeedbackRequestValidationError{
			field:  "UserId",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetFeedback() == nil {
		err := SubmitFeedbackRequestValidationError{
			field:  "Feedback",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetFeedback()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubmitFeedbackRequestValidationError{
					field:  "Feedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubmitFeedbackRequestValidationError{
					field:  "Feedback",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFeedback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitFeedbackRequestValidationError{
				field:  "Feedback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetMetadata() == nil {
		err := SubmitFeedbackRequestValidationError{
			field:  "Metadata",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubmitFeedbackRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubmitFeedbackRequestValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitFeedbackRequestValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SubmitFeedbackRequestMultiError(errors)
	}
	return nil
}

// SubmitFeedbackRequestMultiError is an error wrapping multiple validation
// errors returned by SubmitFeedbackRequest.ValidateAll() if the designated
// constraints aren't met.
type SubmitFeedbackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitFeedbackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitFeedbackRequestMultiError) AllErrors() []error { return m }

// SubmitFeedbackRequestValidationError is the validation error returned by
// SubmitFeedbackRequest.Validate if the designated constraints aren't met.
type SubmitFeedbackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitFeedbackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitFeedbackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitFeedbackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitFeedbackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitFeedbackRequestValidationError) ErrorName() string {
	return "SubmitFeedbackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitFeedbackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitFeedbackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitFeedbackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitFeedbackRequestValidationError{}

// Validate checks the field values on SubmitFeedbackResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubmitFeedbackResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubmitFeedbackResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubmitFeedbackResponseMultiError, or nil if none found.
func (m *SubmitFeedbackResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SubmitFeedbackResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SubmitFeedbackResponseMultiError(errors)
	}
	return nil
}

// SubmitFeedbackResponseMultiError is an error wrapping multiple validation
// errors returned by SubmitFeedbackResponse.ValidateAll() if the designated
// constraints aren't met.
type SubmitFeedbackResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubmitFeedbackResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubmitFeedbackResponseMultiError) AllErrors() []error { return m }

// SubmitFeedbackResponseValidationError is the validation error returned by
// SubmitFeedbackResponse.Validate if the designated constraints aren't met.
type SubmitFeedbackResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitFeedbackResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitFeedbackResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitFeedbackResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitFeedbackResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitFeedbackResponseValidationError) ErrorName() string {
	return "SubmitFeedbackResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitFeedbackResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitFeedbackResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitFeedbackResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitFeedbackResponseValidationError{}
