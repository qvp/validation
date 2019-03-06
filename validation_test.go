package validation

import (
	"reflect"
	"testing"
)

func BenchmarkDate(b *testing.B) {
	value := reflect.ValueOf("01-12-2019")
	params := []interface{}{"02-01-2006"}

	for n := 0; n < b.N; n++ {
		_ = Date(value, params...)
	}
}

func BenchmarkDate3(b *testing.B) {
	value := reflect.ValueOf("01-12-2019")
	params := []interface{}{"02-01-2006"}

	for n := 0; n < b.N; n++ {
		_ = Date3(value, params...)
	}
}
