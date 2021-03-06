// Code generated by protoc-gen-validate
// source: envoy/api/v2/core/config_source.proto
// DO NOT EDIT!!!

package core

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on ApiConfigSource with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ApiConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := ApiConfigSource_ApiType_name[int32(m.GetApiType())]; !ok {
		return ApiConfigSourceValidationError{
			Field:  "ApiType",
			Reason: "value must be one of the defined enum values",
		}
	}

	for idx, item := range m.GetGrpcServices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return ApiConfigSourceValidationError{
					Field:  fmt.Sprintf("GrpcServices[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetRefreshDelay()).(interface {
		Validate() error
	}); ok {
		if err := v.Validate(); err != nil {
			return ApiConfigSourceValidationError{
				Field:  "RefreshDelay",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if d := m.GetRequestTimeout(); d != nil {
		dur := *d

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return ApiConfigSourceValidationError{
				Field:  "RequestTimeout",
				Reason: "value must be greater than 0s",
			}
		}

	}

	return nil
}

// ApiConfigSourceValidationError is the validation error returned by
// ApiConfigSource.Validate if the designated constraints aren't met.
type ApiConfigSourceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ApiConfigSourceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApiConfigSource.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ApiConfigSourceValidationError{}

// Validate checks the field values on AggregatedConfigSource with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AggregatedConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AggregatedConfigSourceValidationError is the validation error returned by
// AggregatedConfigSource.Validate if the designated constraints aren't met.
type AggregatedConfigSourceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AggregatedConfigSourceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAggregatedConfigSource.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AggregatedConfigSourceValidationError{}

// Validate checks the field values on ConfigSource with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConfigSource) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ConfigSourceSpecifier.(type) {

	case *ConfigSource_Path:
		// no validation rules for Path

	case *ConfigSource_ApiConfigSource:

		if v, ok := interface{}(m.GetApiConfigSource()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					Field:  "ApiConfigSource",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	case *ConfigSource_Ads:

		if v, ok := interface{}(m.GetAds()).(interface {
			Validate() error
		}); ok {
			if err := v.Validate(); err != nil {
				return ConfigSourceValidationError{
					Field:  "Ads",
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	default:
		return ConfigSourceValidationError{
			Field:  "ConfigSourceSpecifier",
			Reason: "value is required",
		}

	}

	return nil
}

// ConfigSourceValidationError is the validation error returned by
// ConfigSource.Validate if the designated constraints aren't met.
type ConfigSourceValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ConfigSourceValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfigSource.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ConfigSourceValidationError{}
