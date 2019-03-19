package validation

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type stringT string
type intT int
type uint8T uint8

// For test values without parameters
type testValues map[interface{}]bool

// For test values with parameters
type testValuesP []struct {
	Value  interface{}
	Params []interface{}
	Valid  bool
}

func runValidatorP(t *testing.T, fn ValidatorP, items testValuesP) {
	for _, item := range items {
		err := fn(reflect.ValueOf(item.Value), item.Params...)
		if (err != nil && item.Valid) || (err == nil && item.Valid == false) {
			t.Error(fmt.Sprintf("on value «%s» params: «%s»", item.Value, item.Params))
		}
	}
}

func runValidator(t *testing.T, fn Validator, items testValues) {
	for value, valid := range items {
		err := fn(reflect.ValueOf(value))
		if (err != nil && valid) || (err == nil && valid == false) {
			t.Error(fmt.Sprintf("on value «%s»", value))
		}
	}
}

func TestEmpty(t *testing.T) {
	var itf interface{}

	var items = testValues{
		"":    true,
		0:     true,
		0.0:   true,
		itf:   true,
		" ":   false,
		"abc": false,
	}

	runValidator(t, Empty, items)
}

func TestEmail(t *testing.T) {
	var items = testValues{
		"mail@example.com":          true,
		"m@d.io":                    true,
		"123mail@mail.com":          true,
		"m-ail_com@mail-server.com": true,
		"sdfsf@fsfs":                false,
		"asdad":                     false,
		"":                          false,
	}

	runValidator(t, Email, items)
}

func TestURL(t *testing.T) {
	var items = testValues{
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

func TestAccepted(t *testing.T) {
	var items = testValues{
		"yes":  true,
		"y":    true,
		"on":   true,
		"1":    true,
		"true": true,
		1:      true,
		"YeS":  false,
		"off":  false,
	}

	runValidator(t, Accepted, items)
}

func TestMin(t *testing.T) {
	var items = testValuesP{
		{"abc", []interface{}{"3"}, true},
		{-4, []interface{}{"-9"}, true},
		{stringT("string str"), []interface{}{"3"}, true},
		{9, []interface{}{"18"}, false},
		{uint8T(9), []interface{}{"18"}, false},
		{"Привет", []interface{}{"30"}, false},
	}

	runValidatorP(t, Min, items)
}

func TestMax(t *testing.T) {
	var items = testValuesP{
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

func TestLen(t *testing.T) {
	var items = testValuesP{
		{"ABC", []interface{}{3}, true},
		{[]int{1, 2}, []interface{}{"2"}, true},
		{[2]int{1, 2}, []interface{}{9}, false},
	}

	runValidatorP(t, Len, items)
}

func TestIn(t *testing.T) {
	var items = testValuesP{
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

func TestNotIn(t *testing.T) {
	var items = testValuesP{
		{"A", []interface{}{"A", "B"}, false},
		{1, []interface{}{1, 2}, false},
		{3, []interface{}{1, 2}, true},
		{3, []interface{}{1, 2}, true},
	}

	runValidatorP(t, NotIn, items)
}

func TestHasPrefix(t *testing.T) {
	var items = testValuesP{
		{"prepare", []interface{}{"pre"}, true},
		{"prepare", []interface{}{"PRE"}, false},
		{"prepare", []interface{}{"pres"}, false},
	}

	runValidatorP(t, HasPrefix, items)
}

func TestHasSuffix(t *testing.T) {
	var items = testValuesP{
		{"prepare", []interface{}{"are"}, true},
		{"prepare", []interface{}{"ARE"}, false},
		{"prepare", []interface{}{"ares"}, false},
	}

	runValidatorP(t, HasSuffix, items)
}

func TestRegex(t *testing.T) {
	var items = testValuesP{
		{"123", []interface{}{"^[\\d]{3}$"}, true},
		{"YES", []interface{}{"(?i)yes"}, true},
		{"YES", []interface{}{"yes"}, false},
	}

	runValidatorP(t, Regex, items)
}

func TestAlpha(t *testing.T) {
	var items = testValues{
		"abc":     true,
		"abcXYz":  true,
		"abcАБВ":  false,
		"":        false,
		"123abc":  false,
		"  abc  ": false,
		"a b c":   false,
		"123":     false,
	}

	runValidator(t, Alpha, items)
}

func TestAlphaNumeric(t *testing.T) {
	var items = testValues{
		"abc":     true,
		"abc123":  true,
		"123":     true,
		"":        false,
		"abc234/": false,
	}

	runValidator(t, AlphaNumeric, items)
}

func TestAlphaUnder(t *testing.T) {
	var items = testValues{
		"abc_xyz": true,
		"abc":     true,
		"_":       true,
		"__abc__": true,
		"":        false,
		"abc_234": false,
		"435":     false,
		"--abc_":  false,
	}

	runValidator(t, AlphaUnder, items)
}

func TestAlphaDash(t *testing.T) {
	var items = testValues{
		"abc-xyz": true,
		"abc":     true,
		"-":       true,
		"--abc--": true,
		"":        false,
		"abc_234": false,
		"435":     false,
		"--abc_":  false,
	}

	runValidator(t, AlphaDash, items)
}

func TestASCII(t *testing.T) {
	var items = testValues{
		" !\"#$%&\\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~": true,
		"АБВ": false,
		"":    false,
		"호랑이": false,
	}

	runValidator(t, ASCII, items)
}

func TestInt(t *testing.T) {
	var items = testValues{
		"123":   true,
		"0":     true,
		"0.123": false,
		"":      false,
		"  1.5": false,
		"1,5":   false,
	}

	runValidator(t, Int, items)
}

func TestFloat(t *testing.T) {
	var items = testValues{
		"0.123": true,
		"123":   false,
		"":      false,
		"  1.5": false,
		"1,5":   false,
	}

	runValidator(t, Float, items)
}

func TestJSON(t *testing.T) {
	var items = testValues{
		"\"string value\"": true,
		"[1,4]":            true,
		"[]":               true,
		"{}":               true,
		"{\"id\":5}":       true,
		"{'id':5}":         false,
		"string":           false,
		"{\"string\"}":     false,
	}

	runValidator(t, JSON, items)
}

func TestIp(t *testing.T) {
	var items = testValues{
		"127.0.0.1":                true,
		"64:ff9b::255.255.255.255": true,
		"400.0.0.0":                false,
		"192.168.0.0/16":           false,
		"":                         false,
		"123.456":                  false,
	}

	runValidator(t, Ip, items)
}

func TestIpv4(t *testing.T) {
	var items = testValues{
		"216.3.128.12":             true,
		"127.0.0.1":                true,
		"192.168.0.0":              true,
		"64:ff9b::255.255.255.255": false,
		"192.168.0.0/16":           false,
		"":                         false,
		"123.456":                  false,
	}

	runValidator(t, Ipv4, items)
}

func TestIpv6(t *testing.T) {
	var items = testValues{
		"64:ff9b::255.255.255.255":                true,
		"FE80:0000:0000:0000:0202:B3FF:FE1E:8329": true,
		"[2001:db8:0:1]:80":                       false,
		"216.3.128.12":                            false,
		"127.0.0.1":                               false,
		"192.168.0.0":                             false,
		"192.168.0.0/16":                          false,
		"":                                        false,
		"123.456":                                 false,
	}

	runValidator(t, Ipv6, items)
}

func TestContains(t *testing.T) {
	var items = testValuesP{
		{"Prepare", []interface{}{"rep"}, true},
		{"Prepare", []interface{}{"Pre"}, true},
		{"Prepare", []interface{}{"are"}, true},
		{"Prepare", []interface{}{""}, true},
		{"", []interface{}{""}, true},
		{"Prepare", []interface{}{"res"}, false},
		{"Prepare", []interface{}{"aRe"}, false},
		{"Prepare", []interface{}{"pre"}, false},
	}

	runValidatorP(t, Contains, items)
}

func TestGt(t *testing.T) {
	var items = testValuesP{
		{"ABC", []interface{}{2}, true},
		{10, []interface{}{2}, true},
		{[]int{1, 2, 3}, []interface{}{"2"}, true},
		{"ABC", []interface{}{3}, false},
		{"ABC", []interface{}{5}, false},
		{"", []interface{}{2}, false},
		{5, []interface{}{20}, false},
	}

	runValidatorP(t, Gt, items)
}

func TestLt(t *testing.T) {
	var items = testValuesP{
		{"ABC", []interface{}{4}, true},
		{10, []interface{}{40}, true},
		{[]int{1, 2, 3}, []interface{}{"4"}, true},
		{"", []interface{}{2}, true},
		{"ABC", []interface{}{3}, false},
		{"ABC", []interface{}{1}, false},
		{50, []interface{}{20}, false},
	}

	runValidatorP(t, Lt, items)
}

func TestHasKeys(t *testing.T) {

	m := map[string]bool{
		"x": true,
		"y": false,
		"z": true,
	}

	var items = testValuesP{
		{m, []interface{}{"x", "y"}, true},
		{m, []interface{}{"x", "y", "f"}, false},
		{m, []interface{}{""}, false},
	}

	runValidatorP(t, HasKeys, items)
}

func TestHasOnlyKeys(t *testing.T) {

	m := map[string]bool{
		"x": true,
		"y": false,
		"z": true,
	}

	var items = testValuesP{
		{m, []interface{}{"x", "y", "z"}, true},
		{m, []interface{}{"x", "y"}, false},
		{m, []interface{}{"x", "y", "f"}, false},
		{m, []interface{}{""}, false},
	}

	runValidatorP(t, HasOnlyKeys, items)
}

func TestTime(t *testing.T) {
	var items = testValues{
		"12:30:01": true,
		"00:00:00": true,
		"00:00:90": false,
		"":         false,
		"abc":      false,
		"12:00":    false,
	}

	runValidator(t, Time, items)
}

func TestUpperCase(t *testing.T) {
	var items = testValues{
		"ABC":       true,
		"":          true,
		"АБВ":       true,
		" ":         true,
		"13 - 54.5": true,
		"ABC ":      true,
		"abc":       false,
		"abcABC":    false,
	}

	runValidator(t, UpperCase, items)
}

func TestLowerCase(t *testing.T) {
	var items = testValues{
		"abc":       true,
		"":          true,
		"абв":       true,
		" ":         true,
		"13 - 54.5": true,
		"abc ":      true,
		"ABC":       false,
		"abcABC":    false,
	}

	runValidator(t, LowerCase, items)
}

func TestPassword(t *testing.T) {
	var items = testValues{
		"abcABC0123":          true,
		"abcABC0123#24!!!":    true,
		"abcA01":              false,
		"abczyxfsdfjsf":       false,
		"abczyxfsdfjsf123123": false,
		"SDFSFKLSFLSF123123":  false,
	}

	runValidator(t, Password, items)
}

func TestDate(t *testing.T) {
	var items = testValuesP{
		{"01-12-2019", []interface{}{"02-01-2006"}, true},
		{"01-52-2019", []interface{}{"02-01-2006"}, false},
		{"fake str", []interface{}{"02-01-2006"}, false},
	}

	runValidatorP(t, Date, items)
}

func TestDateGte(t *testing.T) {
	now := time.Now()

	var items = testValuesP{
		{now.Format("02-01-2006"), []interface{}{"02-01-2006", "-1D"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "01-11-2017"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "01-12-2019"}, true},
		{now.Format("02-01-2006"), []interface{}{"02-01-2006", "+1D"}, false},
		{"01-12-2019", []interface{}{"02-01-2006", "02-12-2019"}, false},
		{"01-52-2019", []interface{}{"02-01-2006", "02-12-2019"}, false},
		{"fake str", []interface{}{"02-01-2006", "02-12-2019"}, false},
	}

	runValidatorP(t, DateGte, items)
}

func TestDateLte(t *testing.T) {
	now := time.Now()

	var items = testValuesP{
		{now.Format("02-01-2006"), []interface{}{"02-01-2006", "+1D"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "10-12-2019"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "01-12-2019"}, true},
		{now.Format("02-01-2006"), []interface{}{"02-01-2006", "-1D"}, false},
		{"01-12-2019", []interface{}{"02-01-2006", "02-11-2019"}, false},
		{"01-52-2019", []interface{}{"02-01-2006", "02-12-2019"}, false},
		{"fake str", []interface{}{"02-01-2006", "02-12-2019"}, false},
	}

	runValidatorP(t, DateLte, items)
}

func TestDateGt(t *testing.T) {
	now := time.Now()

	var items = testValuesP{
		{now.Format("02-01-2006"), []interface{}{"02-01-2006", "-1D"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "10-12-2010"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "01-12-2019"}, false},
		{"01-52-2019", []interface{}{"02-01-2006", "02-12-2019"}, false},
		{"fake str", []interface{}{"02-01-2006", "02-12-2019"}, false},
	}

	runValidatorP(t, DateGt, items)
}

func TestDateLt(t *testing.T) {
	now := time.Now()

	var items = testValuesP{
		{now.Format("02-01-2006"), []interface{}{"02-01-2006", "+1D"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "10-12-2020"}, true},
		{"01-12-2019", []interface{}{"02-01-2006", "01-12-2019"}, false},
		{"01-52-2019", []interface{}{"02-01-2006", "02-12-2019"}, false},
		{"fake str", []interface{}{"02-01-2006", "02-12-2019"}, false},
	}

	runValidatorP(t, DateLt, items)
}

func TestCountryCode2(t *testing.T) {
	var items = testValues{
		"RU":  true,
		"ZU":  false,
		"ru":  false,
		"RUS": false,
	}

	runValidator(t, CountryCode2, items)
}

func TestCountryCode3(t *testing.T) {
	var items = testValues{
		"AFG": true,
		"RU":  false,
		"ZU":  false,
		"rus": false,
		"ZUZ": false,
	}

	runValidator(t, CountryCode3, items)
}

func TestCurrencyCode(t *testing.T) {
	var items = testValues{
		"RUB": true,
		"rub": false,
		"RU":  false,
	}

	runValidator(t, CurrencyCode, items)
}

func TestLanguageCode2(t *testing.T) {
	var items = testValues{
		"ru":  true,
		"zu":  true,
		"RU":  false,
		"rus": false,
		"ZUZ": false,
	}

	runValidator(t, LanguageCode2, items)
}

func TestLanguageCode3(t *testing.T) {
	var items = testValues{
		"rus": true,
		"RU":  false,
		"ru":  false,
	}

	runValidator(t, LanguageCode3, items)
}

func TestCreditCard(t *testing.T) {
	var items = testValues{
		"4111111111111111":           true,  // Visa
		"5500000000000004":           true,  // MasterCard
		"340000000000009":            true,  // American Express
		"30000000000004":             true,  // Diner's Club
		"6011000000000004":           true,  // Discover
		"201400000000009":            true,  // en Route
		"3088000000000009":           true,  // JCB
		"     4111111111111111     ": false, // Visa
		"0000 1111 1111 1111":        false,
		"4111 1111 1111 1111":        false,
		"4111 1111 1111":             false,
		"":                           false,
		"4111 1111 1111 1111 1111 1111 1111 1111": false,
	}

	runValidator(t, CreditCard, items)
}
