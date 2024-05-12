// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: promotion/service.proto

package promotion

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

// Validate checks the field values on JoinCampaignRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *JoinCampaignRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JoinCampaignRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// JoinCampaignRequestMultiError, or nil if none found.
func (m *JoinCampaignRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *JoinCampaignRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CampaignId

	if len(errors) > 0 {
		return JoinCampaignRequestMultiError(errors)
	}

	return nil
}

// JoinCampaignRequestMultiError is an error wrapping multiple validation
// errors returned by JoinCampaignRequest.ValidateAll() if the designated
// constraints aren't met.
type JoinCampaignRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JoinCampaignRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JoinCampaignRequestMultiError) AllErrors() []error { return m }

// JoinCampaignRequestValidationError is the validation error returned by
// JoinCampaignRequest.Validate if the designated constraints aren't met.
type JoinCampaignRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinCampaignRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinCampaignRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinCampaignRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinCampaignRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinCampaignRequestValidationError) ErrorName() string {
	return "JoinCampaignRequestValidationError"
}

// Error satisfies the builtin error interface
func (e JoinCampaignRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinCampaignRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinCampaignRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinCampaignRequestValidationError{}

// Validate checks the field values on JoinCampaignReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *JoinCampaignReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on JoinCampaignReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// JoinCampaignReplyMultiError, or nil if none found.
func (m *JoinCampaignReply) ValidateAll() error {
	return m.validate(true)
}

func (m *JoinCampaignReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Message

	if len(errors) > 0 {
		return JoinCampaignReplyMultiError(errors)
	}

	return nil
}

// JoinCampaignReplyMultiError is an error wrapping multiple validation errors
// returned by JoinCampaignReply.ValidateAll() if the designated constraints
// aren't met.
type JoinCampaignReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m JoinCampaignReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m JoinCampaignReplyMultiError) AllErrors() []error { return m }

// JoinCampaignReplyValidationError is the validation error returned by
// JoinCampaignReply.Validate if the designated constraints aren't met.
type JoinCampaignReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinCampaignReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinCampaignReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinCampaignReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinCampaignReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinCampaignReplyValidationError) ErrorName() string {
	return "JoinCampaignReplyValidationError"
}

// Error satisfies the builtin error interface
func (e JoinCampaignReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinCampaignReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinCampaignReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinCampaignReplyValidationError{}
