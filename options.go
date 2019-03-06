package validation

import (
	"reflect"
)

// Value must be not empty
// Len must be greater than zero for String, Array, Slice, Map
// Value must be not nil for Pointer, Interface
// Integer, float values must be greater than zero
func Required(value reflect.Value) error {
	return nil
}

// If this option is present validation will not be performed for this field
func Ignore(value reflect.Value) error {
	return nil
}

// If this option is present validation will be performed before the first error
func Lazy(value reflect.Value) error {
	return nil
}
