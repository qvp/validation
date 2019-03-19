package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
)

var (
	regexAlpha        = regexp.MustCompile("^[a-zA-Z]+$")
	regexAlphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
	regexAlphaUnder   = regexp.MustCompile("^[a-zA-Z_]+$")
	regexAlphaDash    = regexp.MustCompile("^[a-zA-Z-]+$")
	regexInt          = regexp.MustCompile("^[-+]?[0-9]+$")
	regexFloat        = regexp.MustCompile("^[0-9]+\\.[0-9]+$")
	regexEmail        = regexp.MustCompile("(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\\])")

	errorWrongType = "Value type not allowed."
)

// Value must be empty
// Len must zero for String, Array, Slice, Map
// Value must be not nil for Pointer, Interface
// Integer, float values must be greater than zero
// Value kind: String, Array, Slice, Map, Number types, Interface, Pointer
// It ignore another types
func Empty(value reflect.Value) error {
	switch value.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if value.Len() != 0 {
			return errorMessage("empty")
		}
	case reflect.Ptr, reflect.Interface:
		if !value.IsNil() {
			return errorMessage("empty")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		size := size(value)
		if size != 0 {
			return errorMessage("empty")
		}
	}
	return nil
}

// Value must be a valid email address
// Value kind: String
// It panics if another types given
func Email(value reflect.Value) error {
	return regexValidator(regexEmail, "email", value)
}

// Value must be a valid URL address
// Value kind: String
// It panics if another types given
func URL(value reflect.Value) error {
	u, err := url.ParseRequestURI(value.String())
	if err != nil || len(u.Host) == 0 {
		return errors.New(Messages["url"])
	}
	return nil
}

// Value must be in "yes", "on", "1", "y", "true"
// Value kind: String
// It panics if another types given
func Accepted(value reflect.Value) error {
	params := []interface{}{"yes", "on", "1", "y", "true"}
	if err := In(value, params...); err != nil {
		return errorMessage("accepted")
	}
	return nil
}

// Value must be greater or equal than min
// Value kind: String, Array, Slice, Map, Number types
// It panics if another types given
func Min(value reflect.Value, params ...interface{}) error {
	min, _ := parseFloat(params[0])
	val := size(value)
	if val < min {
		return errorMessage("min", params...)
	}
	return nil
}

// Value must be less or equal than max
// Value kind: String, Array, Slice, Map, Number types
// It panics if another types given
func Max(value reflect.Value, params ...interface{}) error {
	max, _ := parseFloat(params[0])
	val := size(value)
	if val > max {
		return errorMessage("max", params...)
	}
	return nil
}

// Value must has specified length
// Value kind: String, Array, Slice, Map
// It panics if another types given
func Len(value reflect.Value, params ...interface{}) error {
	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		length, _ := parseFloat(params[0])
		if float64(value.Len()) != length {
			return errorMessage("len", params...)
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Value must be in the specified list
// Value kind: String, Number types
// It panics if another types given
func In(value reflect.Value, params ...interface{}) error {
	switch value.Kind() {
	case reflect.String:
		v := value.String()
		for _, item := range params {
			p := fmt.Sprintf("%v", item)
			if v == p {
				return nil
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		v, _ := parseFloat(value)
		for _, item := range params {
			p, _ := parseFloat(item)
			if v == p {
				return nil
			}
		}
	default:
		panic(errorWrongType)
	}
	return errorMessage("in", params...)
}

// Value must not be in specified list
// Value kind: String, Number types
// It panics if another types given
func NotIn(value reflect.Value, params ...interface{}) error {
	if err := In(value, params...); err == nil {
		return errorMessage("not_in", params...)
	}
	return nil
}

// String must be begins with specified string
// Value kind: String
// It panics if another types given
func HasPrefix(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return strings.HasPrefix(value.String(), params[0].(string))
	}
	return stringValidatorP("has_prefix", value, params, fn)
}

// Value must be ends with specified string
// Value kind: String
// It panics if another types given
func HasSuffix(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return strings.HasSuffix(value.String(), params[0].(string))
	}
	return stringValidatorP("has_suffix", value, params, fn)
}

// Value must be matching a specified pattern.
// Value kind: String
// It panics if another types given
func Regex(value reflect.Value, params ...interface{}) error {
	regex, err := regexp.Compile(params[0].(string))
	if err != nil {
		panic(err)
	}
	return regexValidator(regex, "regex", value)
}

// Value must be contains only english letters (pattern: ^[a-zA-Z]+$).
// Value kind: String
// It panics if another types given
func Alpha(value reflect.Value) error {
	return regexValidator(regexAlpha, "alpha", value)
}

// Value must be contains only english letters and digits (pattern: ^[a-zA-Z0-9]+$).
// Value kind: String
// It panics if another types given
func AlphaNumeric(value reflect.Value) error {
	return regexValidator(regexAlphaNumeric, "alpha_numeric", value)
}

// Value must be contains only english letters and underscores (pattern: ^[a-zA-Z_]+$).
// Value kind: String
// It panics if another types given
func AlphaUnder(value reflect.Value) error {
	return regexValidator(regexAlphaUnder, "alpha_under", value)
}

// Value must be contains only english letters and dashes (pattern: ^[a-zA-Z-]+$).
// Value kind: String
// It panics if another types given
func AlphaDash(value reflect.Value) error {
	return regexValidator(regexAlphaDash, "alpha_dash", value)
}

// Value must be contain only ASCII characters
// Value kind: String
// It panics if another types given
func ASCII(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		if value.Len() == 0 {
			return false
		}
		s := value.String()
		for i := 0; i < len(s); i++ {
			if s[i] > unicode.MaxASCII {
				return false
			}
		}
		return true
	}
	return stringValidator("ascii", value, fn)
}

// Value must be contains only digits (pattern: ^[-+]?[0-9]+$).
// Value kind: String
// It panics if another types given
func Int(value reflect.Value) error {
	return regexValidator(regexInt, "int", value)
}

// Value must be contains only float number (pattern: ^[0-9]+\.[0-9]+$).
// Value kind: String
// It panics if another types given
func Float(value reflect.Value) error {
	return regexValidator(regexFloat, "float", value)
}

// Value must be a valid JSON.
// Value kind: String
// It panics if another types given
func JSON(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		var j json.RawMessage
		return json.Unmarshal([]byte(value.String()), &j) == nil
	}
	return stringValidator("json", value, fn)
}

// Value must be a valid v4 or v6 IP address
// Value kind: String
// It panics if another types given
func Ip(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		ip := net.ParseIP(value.String())
		return ip != nil
	}
	return stringValidator("ip", value, fn)
}

// Value must be a valid IP v4 address
// Value kind: String
// It panics if another types given
func Ipv4(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		ip := net.ParseIP(value.String())
		return ip != nil && ip.To4() != nil
	}
	return stringValidator("ipv4", value, fn)
}

// Value must a valid Ip v6 address
// Value kind: String
// It panics if another types given
func Ipv6(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		ip := net.ParseIP(value.String())
		return ip != nil && ip.To4() == nil
	}
	return stringValidator("ipv6", value, fn)
}

// Value must be contains specified string
// Value kind: String
// It panics if another types given
func Contains(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return strings.Contains(value.String(), params[0].(string))
	}
	return stringValidatorP("contains", value, params, fn)
}

// Value must be greater than specified
// Numbers are compared by value
// String, Slice, Array, Maps are compared by length
// Value kind: Number, String, Slice, Array, Map
// It panics if another types given
func Gt(value reflect.Value, params ...interface{}) error {
	valueSize := size(value)
	paramSize, _ := parseFloat(params[0])
	if valueSize <= paramSize {
		return errorMessage("gt", params...)
	}
	return nil
}

// Value must be lower than specified
// Numbers are compared by value
// String, Slice, Array, Maps are compared by length
// Value kind: Number, String, Slice, Array, Map
// It panics if another types given
func Lt(value reflect.Value, params ...interface{}) error {
	valueSize := size(value)
	paramSize, _ := parseFloat(params[0])
	if valueSize >= paramSize {
		return errorMessage("lt", params...)
	}
	return nil
}

// Value must have specified keys but not limit to them
// Only string keys supported
// Value kind: Map
// It panics if another types given
func HasKeys(value reflect.Value, params ...interface{}) error {
	switch value.Kind() {
	case reflect.Map:
		if value.Len() < len(params) {
			return errorMessage("has_keys", params...)
		}
		var found bool
		keys := value.MapKeys()
		for _, param := range params {
			found = false
			for _, key := range keys {
				if param.(string) == key.String() {
					found = true
				}
			}
			if !found {
				return errorMessage("has_keys", params...)
			}
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Value must have only specified keys
// Only string keys supported
// Value kind: Map
// It panics if another types given
func HasOnlyKeys(value reflect.Value, params ...interface{}) error {
	switch value.Kind() {
	case reflect.Map:
		if value.Len() != len(params) {
			return errorMessage("has_only_keys", params...)
		}
		if err := HasKeys(value, params...); err != nil {
			return errorMessage("has_only_keys", params...)
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Value must be a valid time in format 15:04:05
// Value kind: String
// It panics if another types given
func Time(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		_, err := time.Parse("15:04:05", value.String())
		return err == nil
	}
	return stringValidator("time", value, fn)
}

// Value must be in upper case
// Value kind: String
// It panics if another types given
func UpperCase(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		s := value.String()
		return s == strings.ToUpper(s)
	}
	return stringValidator("upper_case", value, fn)
}

// Value must be in lower case
// Value kind: String
// It panics if another types given
func LowerCase(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		s := value.String()
		return s == strings.ToLower(s)
	}
	return stringValidator("lower_case", value, fn)
}

// Value must contains at least english letters in both cases, numbers and have minimum length 8
// Value kind: String
// It panics if another types given
func Password(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		if value.Len() < 8 {
			return false
		}

		var a, A, d bool
		for _, r := range value.String() {
			if r >= 'a' && r <= 'z' {
				a = true
			}
			if r >= 'A' && r <= 'Z' {
				A = true
			}
			if r >= '0' && r <= '9' {
				d = true
			}
		}
		return a && A && d
	}
	return stringValidator("password", value, fn)
}

// Value must be fit to the specified layout
func Date(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		_, err := time.Parse(params[0].(string), value.String())
		return err == nil
	}
	return stringValidatorP("date", value, params, fn)
}

// Value must be a valid date and greater or equal specified
func DateGte(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "gte")
	}
	return stringValidatorP("date_gte", value, params, fn)
}

// Value must be a valid date and lower or equal specified
func DateLte(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "lte")
	}
	return stringValidatorP("date_lte", value, params, fn)
}

// Value must be a valid date and greater than specified
func DateGt(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "gt")
	}
	return stringValidatorP("date_gt", value, params, fn)
}

// Value must be a valid date and lower than specified
func DateLt(value reflect.Value, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "lt")
	}
	return stringValidatorP("date_lt", value, params, fn)
}

// Value must be a valid country code in ISO2 format
func CountryCode2(value reflect.Value) error {
	return codeValidator("country_code2", value, 2, CountryCodes2)
}

// Value must be a valid country code in ISO3 format
func CountryCode3(value reflect.Value) error {
	return codeValidator("country_code3", value, 3, CountryCodes3)
}

// Value must be a valid currency code
func CurrencyCode(value reflect.Value) error {
	return codeValidator("currency_code", value, 3, CurrencyCodes)
}

// Value must be a valid language code in ISO2 format
func LanguageCode2(value reflect.Value) error {
	return codeValidator("language_code2", value, 2, LanguageCodes2)
}

// Value must be a valid language code in ISO3 format
func LanguageCode3(value reflect.Value) error {
	return codeValidator("language_code3", value, 3, LanguageCodes3)
}

// Value must be a valid credit card number
// It uses luhn algorithm
func CreditCard(value reflect.Value) error {
	fn := func(value reflect.Value) bool {
		return luhn(value.String())
	}
	return stringValidator("credit_card", value, fn)
}

// Helper for creating validators what check string codes
func codeValidator(ruleName string, value reflect.Value, length int, codes []string) error {
	switch value.Kind() {
	case reflect.String:
		if value.Len() != length {
			return errorMessage(ruleName)
		}
		if ok := in(value.String(), codes); !ok {
			return errorMessage(ruleName)
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Helper for creating validators based on regexp rules
func regexValidator(regex *regexp.Regexp, ruleName string, value reflect.Value) error {
	switch value.Kind() {
	case reflect.String:
		if ok := regex.MatchString(value.String()); !ok {
			return errorMessage(ruleName)
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Helper for creating validators based on strings
func stringValidator(ruleName string, value reflect.Value, fn fnStringValidator) error {
	switch value.Kind() {
	case reflect.String:
		ok := fn(value)
		if !ok {
			return errorMessage(ruleName)
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Helper for creating validators with parameters based on strings
func stringValidatorP(ruleName string, value reflect.Value, params []interface{}, fn fnStringValidatorP) error {
	switch value.Kind() {
	case reflect.String:
		ok := fn(value, params)
		if !ok {
			return errorMessage(ruleName, params...)
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Helper for creating date comparision validators
func dateComparison(value string, params []interface{}, fnName string) bool {
	layout := params[0].(string)
	placeholder := DatePlaceholder(params[1].(string))

	date, err := time.Parse(layout, value)
	if err != nil {
		return false
	}

	comparedDate, err := GetDate(placeholder)
	if err == nil {
		comparedDate, _ = time.Parse(layout, comparedDate.Format(layout))
	} else {
		comparedDate, err = time.Parse(layout, string(placeholder))
		if err != nil {
			return false
		}
	}

	switch fnName {
	case "gte":
		return date.After(comparedDate) || date.Equal(comparedDate)
	case "lte":
		return date.Before(comparedDate) || date.Equal(comparedDate)
	case "gt":
		return date.After(comparedDate)
	case "lt":
		return date.Before(comparedDate)
	default:
		return false
	}
}