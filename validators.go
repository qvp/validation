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

// Custom validation function
type Validator func(interface{}, OptionList, ...interface{}) error

// Custom validation functions map
type ValidatorMap map[string]Validator

// List of build-in validators
var validators = map[string]Validator{
	"empty":          Empty,
	"email":          Email,
	"url":            URL,
	"accepted":       Accepted,
	"alpha":          Alpha,
	"alpha_under":    AlphaUnder,
	"alpha_dash":     AlphaDash,
	"ascii":          ASCII,
	"int":            Int,
	"float":          Float,
	"json":           JSON,
	"ip":             Ip,
	"ipv4":           Ipv4,
	"ipv6":           Ipv6,
	"time":           Time,
	"upper_case":     UpperCase,
	"lower_case":     LowerCase,
	"country_code2":  CountryCode2,
	"country_code3":  CountryCode3,
	"currency_code":  CurrencyCode,
	"language_code2": LanguageCode2,
	"language_code3": LanguageCode3,
	"credit_card":    CreditCard,
	"password":       Password,
	"min":            min,
	"max":            Max,
	"len":            Len,
	"in":             In,
	"not_in":         NotIn,
	"date":           Date,
	"regex":          Regex,
	"contains":       Contains,
	"gt":             Gt,
	"lt":             Lt,
	"date_gte":       DateGte,
	"date_lte":       DateLte,
	"date_gt":        DateGt,
	"date_lt":        DateLt,
	"has_prefix":     HasPrefix,
	"has_suffix":     HasSuffix,
	"has_keys":       HasKeys,
	"has_only_keys":  HasOnlyKeys,
}

// Map of custom validation functions
var Validators = ValidatorMap{}

// Regex validation patterns
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

// Add new validation function
func (v ValidatorMap) Add(name string, validator Validator) {
	v[name] = validator
}

// Check what validator exists
func (v ValidatorMap) Has(name string) bool {
	for n := range v {
		if name == n {
			return true
		}
	}
	for n := range validators {
		if name == n {
			return true
		}
	}
	return false
}

// Value must be empty
// Len must zero for String, Array, Slice, Map
// Value must be not nil for Pointer, Interface
// Integer, float values must be greater than zero
// Value kind: String, Array, Slice, Map, Number types, Interface, Pointer
// It ignore another types
func Empty(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map:
		if val.Len() != 0 {
			return errorMessage("empty")
		}
	case reflect.Ptr, reflect.Interface:
		if !val.IsNil() {
			return errorMessage("empty")
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		size := size(val)
		if size != 0 {
			return errorMessage("empty")
		}
	}
	return nil
}

// Value must be a valid email address
// Value kind: String
// It panics if another types given
func Email(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexEmail, "email", value.(reflect.Value))
}

// Value must be a valid URL address
// Value kind: String
// It panics if another types given
func URL(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.String:
		u, err := url.ParseRequestURI(val.String())
		if err != nil || len(u.Host) == 0 {
			return errors.New(Messages["url"])
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Value must be in "yes", "on", "1", "y", "true"
// Value kind: String
// It panics if another types given
func Accepted(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.String:
		allowed := []interface{}{"yes", "on", "1", "y", "true"}
		if err := In(value, options, allowed...); err != nil {
			return errorMessage("accepted")
		}
	default:
		panic(errorWrongType)
	}
	return nil
}

// Value must be greater or equal than min
// Value kind: String, Array, Slice, Map, Number types
// It panics if another types given
func min(value interface{}, options OptionList, params ...interface{}) error {
	min, _ := parseFloat(params[0])
	val := size(value.(reflect.Value))
	if val < min {
		return errorMessage("min", params...)
	}
	return nil
}

// Value must be less or equal than max
// Value kind: String, Array, Slice, Map, Number types
// It panics if another types given
func Max(value interface{}, options OptionList, params ...interface{}) error {
	max, _ := parseFloat(params[0])
	val := size(value.(reflect.Value))
	if val > max {
		return errorMessage("max", params...)
	}
	return nil
}

// Value must has specified length
// Value kind: String, Array, Slice, Map
// It panics if another types given
func Len(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		length, _ := parseFloat(params[0])
		if float64(val.Len()) != length {
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
func In(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.String:
		v := val.String()
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
func NotIn(value interface{}, options OptionList, params ...interface{}) error {
	if err := In(value, options, params...); err == nil {
		return errorMessage("not_in", params...)
	}
	return nil
}

// String must be begins with specified string
// Value kind: String
// It panics if another types given
func HasPrefix(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return strings.HasPrefix(value.String(), params[0].(string))
	}
	return stringValidator("has_prefix", value.(reflect.Value), params, fn)
}

// Value must be ends with specified string
// Value kind: String
// It panics if another types given
func HasSuffix(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return strings.HasSuffix(value.String(), params[0].(string))
	}
	return stringValidator("has_suffix", value.(reflect.Value), params, fn)
}

// Value must be matching a specified pattern.
// Value kind: String
// It panics if another types given
func Regex(value interface{}, options OptionList, params ...interface{}) error {
	regex, err := regexp.Compile(params[0].(string))
	if err != nil {
		panic(err)
	}
	return regexValidator(regex, "regex", value.(reflect.Value))
}

// Value must be contains only english letters (pattern: ^[a-zA-Z]+$).
// Value kind: String
// It panics if another types given
func Alpha(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexAlpha, "alpha", value.(reflect.Value))
}

// Value must be contains only english letters and digits (pattern: ^[a-zA-Z0-9]+$).
// Value kind: String
// It panics if another types given
func AlphaNumeric(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexAlphaNumeric, "alpha_numeric", value.(reflect.Value))
}

// Value must be contains only english letters and underscores (pattern: ^[a-zA-Z_]+$).
// Value kind: String
// It panics if another types given
func AlphaUnder(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexAlphaUnder, "alpha_under", value.(reflect.Value))
}

// Value must be contains only english letters and dashes (pattern: ^[a-zA-Z-]+$).
// Value kind: String
// It panics if another types given
func AlphaDash(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexAlphaDash, "alpha_dash", value.(reflect.Value))
}

// Value must be contain only ASCII characters
// Value kind: String
// It panics if another types given
func ASCII(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
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
	return stringValidator("ascii", value.(reflect.Value), params, fn)
}

// Value must be contains only digits (pattern: ^[-+]?[0-9]+$).
// Value kind: String
// It panics if another types given
func Int(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexInt, "int", value.(reflect.Value))
}

// Value must be contains only float number (pattern: ^[0-9]+\.[0-9]+$).
// Value kind: String
// It panics if another types given
func Float(value interface{}, options OptionList, params ...interface{}) error {
	return regexValidator(regexFloat, "float", value.(reflect.Value))
}

// Value must be a valid JSON.
// Value kind: String
// It panics if another types given
func JSON(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		var j json.RawMessage
		return json.Unmarshal([]byte(value.String()), &j) == nil
	}
	return stringValidator("json", value.(reflect.Value), params, fn)
}

// Value must be a valid v4 or v6 IP address
// Value kind: String
// It panics if another types given
func Ip(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		ip := net.ParseIP(value.String())
		return ip != nil
	}
	return stringValidator("ip", value.(reflect.Value), params, fn)
}

// Value must be a valid IP v4 address
// Value kind: String
// It panics if another types given
func Ipv4(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		ip := net.ParseIP(value.String())
		return ip != nil && ip.To4() != nil
	}
	return stringValidator("ipv4", value.(reflect.Value), params, fn)
}

// Value must a valid Ip v6 address
// Value kind: String
// It panics if another types given
func Ipv6(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		ip := net.ParseIP(value.String())
		return ip != nil && ip.To4() == nil
	}
	return stringValidator("ipv6", value.(reflect.Value), params, fn)
}

// Value must be contains specified string
// Value kind: String
// It panics if another types given
func Contains(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return strings.Contains(value.String(), params[0].(string))
	}
	return stringValidator("contains", value.(reflect.Value), params, fn)
}

// Value must be greater than specified
// Numbers are compared by value
// String, Slice, Array, Maps are compared by length
// Value kind: Number, String, Slice, Array, Map
// It panics if another types given
func Gt(value interface{}, options OptionList, params ...interface{}) error {
	valueSize := size(value.(reflect.Value))
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
func Lt(value interface{}, options OptionList, params ...interface{}) error {
	valueSize := size(value.(reflect.Value))
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
func HasKeys(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.Map:
		if val.Len() < len(params) {
			return errorMessage("has_keys", params...)
		}
		var found bool
		keys := val.MapKeys()
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
func HasOnlyKeys(value interface{}, options OptionList, params ...interface{}) error {
	switch val := value.(reflect.Value); val.Kind() {
	case reflect.Map:
		if val.Len() != len(params) {
			return errorMessage("has_only_keys", params...)
		}
		if err := HasKeys(value, options, params...); err != nil {
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
func Time(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		_, err := time.Parse("15:04:05", value.String())
		return err == nil
	}
	return stringValidator("time", value.(reflect.Value), params, fn)
}

// Value must be in upper case
// Value kind: String
// It panics if another types given
func UpperCase(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		s := value.String()
		return s == strings.ToUpper(s)
	}
	return stringValidator("upper_case", value.(reflect.Value), params, fn)
}

// Value must be in lower case
// Value kind: String
// It panics if another types given
func LowerCase(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		s := value.String()
		return s == strings.ToLower(s)
	}
	return stringValidator("lower_case", value.(reflect.Value), params, fn)
}

// Value must contains at least english letters in both cases, numbers and have minimum length 8
// Value kind: String
// It panics if another types given
func Password(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
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
	return stringValidator("password", value.(reflect.Value), params, fn)
}

// Value must be fit to the specified layout
func Date(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		_, err := time.Parse(params[0].(string), value.String())
		return err == nil
	}
	return stringValidator("date", value.(reflect.Value), params, fn)
}

// Value must be a valid date and greater or equal specified
func DateGte(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "gte")
	}
	return stringValidator("date_gte", value.(reflect.Value), params, fn)
}

// Value must be a valid date and lower or equal specified
func DateLte(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "lte")
	}
	return stringValidator("date_lte", value.(reflect.Value), params, fn)
}

// Value must be a valid date and greater than specified
func DateGt(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "gt")
	}
	return stringValidator("date_gt", value.(reflect.Value), params, fn)
}

// Value must be a valid date and lower than specified
func DateLt(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return dateComparison(value.String(), params, "lt")
	}
	return stringValidator("date_lt", value.(reflect.Value), params, fn)
}

// Value must be a valid country code in ISO2 format
func CountryCode2(value interface{}, options OptionList, params ...interface{}) error {
	return codeValidator("country_code2", value.(reflect.Value), 2, CountryCodes2)
}

// Value must be a valid country code in ISO3 format
func CountryCode3(value interface{}, options OptionList, params ...interface{}) error {
	return codeValidator("country_code3", value.(reflect.Value), 3, CountryCodes3)
}

// Value must be a valid currency code
func CurrencyCode(value interface{}, options OptionList, params ...interface{}) error {
	return codeValidator("currency_code", value.(reflect.Value), 3, CurrencyCodes)
}

// Value must be a valid language code in ISO2 format
func LanguageCode2(value interface{}, options OptionList, params ...interface{}) error {
	return codeValidator("language_code2", value.(reflect.Value), 2, LanguageCodes2)
}

// Value must be a valid language code in ISO3 format
func LanguageCode3(value interface{}, options OptionList, params ...interface{}) error {
	return codeValidator("language_code3", value.(reflect.Value), 3, LanguageCodes3)
}

// Value must be a valid credit card number
// It uses luhn algorithm
func CreditCard(value interface{}, options OptionList, params ...interface{}) error {
	fn := func(value reflect.Value, params []interface{}) bool {
		return luhn(value.String())
	}
	return stringValidator("credit_card", value.(reflect.Value), params, fn)
}

func FileExists(value interface{}, options OptionList, params ...interface{}) error {
	return nil
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
type stringValidatorFunc func(reflect.Value, []interface{}) bool

func stringValidator(ruleName string, value reflect.Value, params []interface{}, fn stringValidatorFunc) error {
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
