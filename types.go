package validation

import "reflect"

type Rule struct {
	Name   string
	Params []interface{}
}

// Validation function
type Validator func(reflect.Value) error

// Validation function with additional parameters
type ValidatorP func(reflect.Value, ...interface{}) error

// Wrapper useful for pass custom validators with parameters
type ValidatorWrapper struct {
	Function ValidatorP
	Params   []interface{}
}

// Represent struct attribute for validation
type StructField struct {
	Name     string
	Value    reflect.Value
	ValidTag string
}

type DatePlaceholder string

type fnStringValidatorP func(reflect.Value, []interface{}) bool
type fnStringValidator func(reflect.Value) bool
