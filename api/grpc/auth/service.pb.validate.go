// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth/service.proto

package auth

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

// Validate checks the field values on RegisterAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterAccountRequestMultiError, or nil if none found.
func (m *RegisterAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FullName

	// no validation rules for PhoneNumber

	// no validation rules for Email

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Birthday

	if len(errors) > 0 {
		return RegisterAccountRequestMultiError(errors)
	}

	return nil
}

// RegisterAccountRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterAccountRequestMultiError) AllErrors() []error { return m }

// RegisterAccountRequestValidationError is the validation error returned by
// RegisterAccountRequest.Validate if the designated constraints aren't met.
type RegisterAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterAccountRequestValidationError) ErrorName() string {
	return "RegisterAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterAccountRequestValidationError{}

// Validate checks the field values on RegisterAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterAccountReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterAccountReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterAccountReplyMultiError, or nil if none found.
func (m *RegisterAccountReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterAccountReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return RegisterAccountReplyMultiError(errors)
	}

	return nil
}

// RegisterAccountReplyMultiError is an error wrapping multiple validation
// errors returned by RegisterAccountReply.ValidateAll() if the designated
// constraints aren't met.
type RegisterAccountReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterAccountReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterAccountReplyMultiError) AllErrors() []error { return m }

// RegisterAccountReplyValidationError is the validation error returned by
// RegisterAccountReply.Validate if the designated constraints aren't met.
type RegisterAccountReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterAccountReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterAccountReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterAccountReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterAccountReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterAccountReplyValidationError) ErrorName() string {
	return "RegisterAccountReplyValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterAccountReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterAccountReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterAccountReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterAccountReplyValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Identity

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on LoginReply_Data with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginReply_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginReply_DataMultiError, or nil if none found.
func (m *LoginReply_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AccountId

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginReply_DataMultiError(errors)
	}

	return nil
}

// LoginReply_DataMultiError is an error wrapping multiple validation errors
// returned by LoginReply_Data.ValidateAll() if the designated constraints
// aren't met.
type LoginReply_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReply_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReply_DataMultiError) AllErrors() []error { return m }

// LoginReply_DataValidationError is the validation error returned by
// LoginReply_Data.Validate if the designated constraints aren't met.
type LoginReply_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReply_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReply_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReply_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReply_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReply_DataValidationError) ErrorName() string { return "LoginReply_DataValidationError" }

// Error satisfies the builtin error interface
func (e LoginReply_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReply_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReply_DataValidationError{}
