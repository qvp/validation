package validation

import (
	"fmt"
	"reflect"
	"testing"
)

type (
	stringT string
	intT    int
	uint8T  uint8
	item    map[interface{}]bool
	itemP   struct {
		Value  interface{}
		Params []interface{}
		Valid  bool
	}
)

func runValidatorP(t *testing.T, fun ValidatorP, items []itemP) {
	for _, item := range items {
		if err := fun(reflect.ValueOf(item.Value), item.Params...); err != nil && item.Valid {
			t.Error(fmt.Sprintf("Fail. Value: %s Params: %s Error: %s", item.Value, item.Params, err))
		}
	}
}

func runValidator(t *testing.T, fun Validator, items item) {
	for value, valid := range items {
		if err := fun(reflect.ValueOf(value)); err != nil && valid {
			t.Error(fmt.Sprintf("Fail. Value: %s Error: %s", value, err))
		}
	}
}

func TestMin(t *testing.T) {
	var items = []itemP{
		{"abc", []interface{}{"3"}, true},
		{9, []interface{}{"18"}, false},
		{uint8T(9), []interface{}{"18"}, false},
		{-4, []interface{}{"-9"}, true},
		{"Привет", []interface{}{"30"}, false},
		{stringT("string str"), []interface{}{"3"}, false},
	}

	runValidatorP(t, Min, items)
}

func TestMax(t *testing.T) {
	var items = []itemP{
		{"abc", []interface{}{"3"}, true},
		{10, []interface{}{"18"}, true},
		{10, []interface{}{18}, true},
		{intT(10), []interface{}{18}, true},
		{-4, []interface{}{"-9"}, false},
		{"Яблоко", []interface{}{"3"}, false},
		{"Груша", []interface{}{"6"}, true},
		{stringT("string str"), []interface{}{"3"}, false},
	}

	runValidatorP(t, Max, items)
}

func TestIn(t *testing.T) {
	var items = []itemP{
		{"yes", []interface{}{"yes", "no"}, true},
		{"nO", []interface{}{"yes", "no"}, false},
		{-4, []interface{}{"-9", "-4"}, true},
		{"Слива", []interface{}{"слива", "персик", "смородина"}, false},
		{stringT("yes"), []interface{}{"yes", "no"}, true},
		{1, []interface{}{1, 2, 3}, true},
		{2.0, []interface{}{1, 2, 3}, true},
	}

	runValidatorP(t, In, items)
}

func TestURL(t *testing.T) {
	var items = item{
		"https://vk.com":     true,
		"http://example.com": true,
		"https://vk.com/":    true,
		"https://vk.co":      true,
		"vk":                 false,
		"vk.com/fake":        false,
		"vk.com":             false,
		4:                    false,
		uint8T(4):            false,
		"/uri":               false,
	}

	runValidator(t, URL, items)
}

func TestDate(t *testing.T) {
	var items = []itemP{
		{"01-12-2019", []interface{}{"02-01-2006"}, true},
		{"01-52-2019", []interface{}{"02-01-2006"}, false},
		{"fake str", []interface{}{"02-01-2006"}, false},
	}

	runValidatorP(t, Date, items)
}

func TestCountryCode2(t *testing.T) {
	var items = item{
		"RU":  true,
		"ZU":  false,
		"ru":  false,
		"RUS": false,
	}

	runValidator(t, CountryCode2, items)
}

func TestCountryCode3(t *testing.T) {
	var items = item{
		"RU":  false,
		"ZU":  false,
		"rus": false,
		"AFG": true,
		"ZUZ": false,
	}

	runValidator(t, CountryCode3, items)
}

func TestRequired(t *testing.T) {
	var items = item{
		0:         false,
		1:         true,
		2.2:       true,
		"":        false,
		"hello":   true,
		[0]int{}:  false,
		[3]int{}:  true,
		[1]int{5}: true,
	}

	runValidator(t, Required, items)
}

func TestAccepted(t *testing.T) {
	var items = item{
		"YeS":  false,
		"yes":  true,
		"y":    true,
		"on":   true,
		"off":  false,
		"1":    true,
		"true": true,
		1:      false,
	}

	runValidator(t, Accepted, items)
}

func TestLen(t *testing.T) {
	var items = []itemP{
		{"ABC", []interface{}{3}, true},
		{[]int{1, 2}, []interface{}{"2"}, true},
		{[2]int{1, 2}, []interface{}{9}, false},
	}

	runValidatorP(t, Len, items)
}

func TestNotIn(t *testing.T) {
	var items = []itemP{
		{"A", []interface{}{"A", "B"}, false},
		{1, []interface{}{1, 2}, false},
		{3, []interface{}{1, 2}, true},
		{3, []interface{}{1, 2}, true},
	}

	runValidatorP(t, NotIn, items)
}

func TestRegex(t *testing.T) {
	var items = []itemP{
		{"123", []interface{}{"^[\\d]{3}$"}, true},
		{"YES", []interface{}{"(?i)yes"}, true},
		{"YES", []interface{}{"yes"}, false},
	}

	runValidatorP(t, Regex, items)
}

func TestBegins(t *testing.T) {
	var items = []itemP{
		{"prepare", []interface{}{"pre"}, true},
		{"prepare", []interface{}{"PRE"}, false},
		{"prepare", []interface{}{"pres"}, false},
	}

	runValidatorP(t, Begins, items)
}

func TestEnds(t *testing.T) {
	var items = []itemP{
		{"prepare", []interface{}{"are"}, true},
		{"prepare", []interface{}{"ARE"}, false},
		{"prepare", []interface{}{"ares"}, false},
	}

	runValidatorP(t, Ends, items)
}
