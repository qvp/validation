package validation

import "reflect"

// Represent struct attribute for validation
type StructField struct {
	Name     string
	Value    reflect.Value
	ValidTag string
}

type DatePlaceholder string
