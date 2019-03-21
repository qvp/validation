package validation

import (
	"fmt"
	"reflect"
	"strings"
)

// Action function type
type Action func(interface{}) interface{}

// List of actions
type ActionMap map[string]Action

// List of present actions
var Actions = ActionMap{
	"trim":  Trim,
	"lower": Lower,
	"upper": Upper,
	"clear": Clear,
}

// Check what action exists
func (a ActionMap) Has(name string) bool {
	_, ok := Actions[name]
	return ok
}

// Add new action
func (a ActionMap) Add(name string, action Action) {
	a[name] = action
}

// Trim spaces before and after string
func Trim(value interface{}) interface{} {
	return strings.TrimSpace(value.(string))
}

// Converts the string to lower case
func Lower(value interface{}) interface{} {
	return strings.ToLower(value.(string))
}

// Converts the string to upper case
func Upper(value interface{}) interface{} {
	return strings.ToUpper(value.(string))
}

// Clear value by set default value
// String transformed to empty string
// Numbers transformed to zero value
// Another types ignored
func Clear(value interface{}) interface{} { //todo maybe add Slice, Array, Map?
	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.String:
		return ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return 0
	default:
		return value
	}
}

//Set default value if value is empty
//func Default(value interface{}) interface{} { //todo this need action options!
//	valueOf := reflect.ValueOf(value)
//	if Empty(valueOf, OptionList{}) == nil {
//		switch valueOf.Kind() {
//		case reflect.:
//
//		}
//	}
//	return value
//}

// Generate key for action
func (a ActionMap) key() string {
	return fmt.Sprintf("_%d_", len(a))
}
