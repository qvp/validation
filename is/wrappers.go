package is

import (
	"github.com/qvp/validation"
	"reflect"
)

//todo return Validator type

func URL(value reflect.Value) error {
	return validation.Validator(validation.URL)(value)
}

func Email(value reflect.Value) error {
	return validation.Email(value)
}

func Required(value reflect.Value) error {
	return validation.Required(value)
}

func Accepted(value reflect.Value) error {
	return validation.Accepted(value)
}

func Alpha(value reflect.Value) error {
	return validation.Alpha(value)
}

func Numeric(value reflect.Value) error {
	return validation.Numeric(value)
}

func AlphaNumeric(value reflect.Value) error {
	return validation.AlphaNumeric(value)
}

func UTF(value reflect.Value) error {
	return validation.UTF(value)
}

func ASCII(value reflect.Value) error {
	return validation.ASCII(value)
}

func Int(value reflect.Value) error {
	return validation.Int(value)
}

func Float(value reflect.Value) error {
	return validation.Float(value)
}

func JSON(value reflect.Value) error {
	return validation.JSON(value)
}

func Ip(value reflect.Value) error {
	return validation.Ip(value)
}

func Ipv4(value reflect.Value) error {
	return validation.Ipv4(value)
}

func Ipv6(value reflect.Value) error {
	return validation.Ipv6(value)
}

func Time(value reflect.Value) error {
	return validation.Time(value)
}

func UpperCase(value reflect.Value) error {
	return validation.UpperCase(value)
}

func LowerCase(value reflect.Value) error {
	return validation.LowerCase(value)
}

func NonZero(value reflect.Value) error {
	return validation.NonZero(value)
}
