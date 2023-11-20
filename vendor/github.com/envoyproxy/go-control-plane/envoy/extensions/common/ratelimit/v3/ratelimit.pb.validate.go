// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/ratelimit/v3/ratelimit.proto

package ratelimitv3

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

	v3 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
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

	_ = v3.RateLimitUnit(0)
)

// Validate checks the field values on RateLimitDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitDescriptor) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitDescriptorMultiError, or nil if none found.
func (m *RateLimitDescriptor) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitDescriptor) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetEntries()) < 1 {
		err := RateLimitDescriptorValidationError{
			field:  "Entries",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RateLimitDescriptorValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RateLimitDescriptorValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RateLimitDescriptorValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetLimit()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitDescriptorValidationError{
					field:  "Limit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitDescriptorValidationError{
					field:  "Limit",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLimit()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitDescriptorValidationError{
				field:  "Limit",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RateLimitDescriptorMultiError(errors)
	}

	return nil
}

// RateLimitDescriptorMultiError is an error wrapping multiple validation
// errors returned by RateLimitDescriptor.ValidateAll() if the designated
// constraints aren't met.
type RateLimitDescriptorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitDescriptorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitDescriptorMultiError) AllErrors() []error { return m }

// RateLimitDescriptorValidationError is the validation error returned by
// RateLimitDescriptor.Validate if the designated constraints aren't met.
type RateLimitDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptorValidationError) ErrorName() string {
	return "RateLimitDescriptorValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptorValidationError{}

// Validate checks the field values on LocalRateLimitDescriptor with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *LocalRateLimitDescriptor) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalRateLimitDescriptor with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LocalRateLimitDescriptorMultiError, or nil if none found.
func (m *LocalRateLimitDescriptor) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalRateLimitDescriptor) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetEntries()) < 1 {
		err := LocalRateLimitDescriptorValidationError{
			field:  "Entries",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalRateLimitDescriptorValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalRateLimitDescriptorValidationError{
						field:  fmt.Sprintf("Entries[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalRateLimitDescriptorValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.GetTokenBucket() == nil {
		err := LocalRateLimitDescriptorValidationError{
			field:  "TokenBucket",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTokenBucket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalRateLimitDescriptorValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalRateLimitDescriptorValidationError{
					field:  "TokenBucket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokenBucket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalRateLimitDescriptorValidationError{
				field:  "TokenBucket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LocalRateLimitDescriptorMultiError(errors)
	}

	return nil
}

// LocalRateLimitDescriptorMultiError is an error wrapping multiple validation
// errors returned by LocalRateLimitDescriptor.ValidateAll() if the designated
// constraints aren't met.
type LocalRateLimitDescriptorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalRateLimitDescriptorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalRateLimitDescriptorMultiError) AllErrors() []error { return m }

// LocalRateLimitDescriptorValidationError is the validation error returned by
// LocalRateLimitDescriptor.Validate if the designated constraints aren't met.
type LocalRateLimitDescriptorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalRateLimitDescriptorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalRateLimitDescriptorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalRateLimitDescriptorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalRateLimitDescriptorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalRateLimitDescriptorValidationError) ErrorName() string {
	return "LocalRateLimitDescriptorValidationError"
}

// Error satisfies the builtin error interface
func (e LocalRateLimitDescriptorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalRateLimitDescriptor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalRateLimitDescriptorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalRateLimitDescriptorValidationError{}

// Validate checks the field values on RateLimitDescriptor_Entry with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitDescriptor_Entry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitDescriptor_Entry with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitDescriptor_EntryMultiError, or nil if none found.
func (m *RateLimitDescriptor_Entry) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitDescriptor_Entry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := RateLimitDescriptor_EntryValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetValue()) < 1 {
		err := RateLimitDescriptor_EntryValidationError{
			field:  "Value",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RateLimitDescriptor_EntryMultiError(errors)
	}

	return nil
}

// RateLimitDescriptor_EntryMultiError is an error wrapping multiple validation
// errors returned by RateLimitDescriptor_Entry.ValidateAll() if the
// designated constraints aren't met.
type RateLimitDescriptor_EntryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitDescriptor_EntryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitDescriptor_EntryMultiError) AllErrors() []error { return m }

// RateLimitDescriptor_EntryValidationError is the validation error returned by
// RateLimitDescriptor_Entry.Validate if the designated constraints aren't met.
type RateLimitDescriptor_EntryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptor_EntryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptor_EntryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptor_EntryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptor_EntryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptor_EntryValidationError) ErrorName() string {
	return "RateLimitDescriptor_EntryValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptor_EntryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor_Entry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptor_EntryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptor_EntryValidationError{}

// Validate checks the field values on RateLimitDescriptor_RateLimitOverride
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *RateLimitDescriptor_RateLimitOverride) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitDescriptor_RateLimitOverride
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// RateLimitDescriptor_RateLimitOverrideMultiError, or nil if none found.
func (m *RateLimitDescriptor_RateLimitOverride) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitDescriptor_RateLimitOverride) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RequestsPerUnit

	if _, ok := v3.RateLimitUnit_name[int32(m.GetUnit())]; !ok {
		err := RateLimitDescriptor_RateLimitOverrideValidationError{
			field:  "Unit",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RateLimitDescriptor_RateLimitOverrideMultiError(errors)
	}

	return nil
}

// RateLimitDescriptor_RateLimitOverrideMultiError is an error wrapping
// multiple validation errors returned by
// RateLimitDescriptor_RateLimitOverride.ValidateAll() if the designated
// constraints aren't met.
type RateLimitDescriptor_RateLimitOverrideMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitDescriptor_RateLimitOverrideMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitDescriptor_RateLimitOverrideMultiError) AllErrors() []error { return m }

// RateLimitDescriptor_RateLimitOverrideValidationError is the validation error
// returned by RateLimitDescriptor_RateLimitOverride.Validate if the
// designated constraints aren't met.
type RateLimitDescriptor_RateLimitOverrideValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptor_RateLimitOverrideValidationError) ErrorName() string {
	return "RateLimitDescriptor_RateLimitOverrideValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptor_RateLimitOverrideValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescriptor_RateLimitOverride.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptor_RateLimitOverrideValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptor_RateLimitOverrideValidationError{}
