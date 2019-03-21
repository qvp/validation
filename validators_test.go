package validation

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type testItem struct {
	Value   interface{}
	Options OptionList
	Params  []interface{}
	IsValid bool
}

// For tests with custom types
type stringT string
type intT int
type uint8T uint8

func TestEmpty(t *testing.T) {
	var itf interface{}

	var items = []testItem{
		{Value: "", IsValid: true},
		{Value: 0, IsValid: true},
		{Value: 0.0, IsValid: true},
		{Value: itf, IsValid: true},
		{Value: " ", IsValid: false},
		{Value: "abc", IsValid: false},
	}

	testItems(t, Empty, items)
}

func TestEmail(t *testing.T) {
	var items = []testItem{
		{Value: "mail@example.com", IsValid: true},
		{Value: "m@d.io", IsValid: true},
		{Value: "123mail@mail.com", IsValid: true},
		{Value: "m-ail_com@mail-server.com", IsValid: true},
		{Value: "sdfsf@fsfs", IsValid: false},
		{Value: "abc", IsValid: false},
		{Value: "", IsValid: false},
	}

	testItems(t, Email, items)
}

func TestURL(t *testing.T) {
	var items = []testItem{
		{Value: "https://vk.com", IsValid: true},
		{Value: "https://vk.com/fake", IsValid: true},
		{Value: "http://example.com", IsValid: true},
		{Value: "https://vk.com/", IsValid: true},
		{Value: "https://vk.co", IsValid: true},
		{Value: "vk.com", IsValid: false},
		{Value: "vk", IsValid: false},
		{Value: "/uri", IsValid: false},
		{Value: "", IsValid: false},
	}

	testItems(t, URL, items)
}

func TestAccepted(t *testing.T) {
	var items = []testItem{
		{Value: "yes", IsValid: true},
		{Value: "y", IsValid: true},
		{Value: "on", IsValid: true},
		{Value: "1", IsValid: true},
		{Value: "true", IsValid: true},
		{Value: "YeS", IsValid: false},
		{Value: "off", IsValid: false},
	}

	testItems(t, Accepted, items)
}

func TestMin(t *testing.T) {
	var items = []testItem{
		{Value: "abc", Params: []interface{}{"3"}, IsValid: true},
		{Value: -4, Params: []interface{}{"-9"}, IsValid: true},
		{Value: stringT("string str"), Params: []interface{}{"3"}, IsValid: true},
		{Value: 9, Params: []interface{}{"18"}, IsValid: false},
		{Value: uint8T(9), Params: []interface{}{"18"}, IsValid: false},
		{Value: "Привет", Params: []interface{}{"30"}, IsValid: false},
	}

	testItems(t, min, items)
}

func TestMax(t *testing.T) {
	var items = []testItem{
		{Value: "abc", Params: []interface{}{"3"}, IsValid: true},
		{Value: 10, Params: []interface{}{"18"}, IsValid: true},
		{Value: 10, Params: []interface{}{18}, IsValid: true},
		{Value: intT(10), Params: []interface{}{18}, IsValid: true},
		{Value: -4, Params: []interface{}{"-9"}, IsValid: false},
		{Value: "Яблоко", Params: []interface{}{"3"}, IsValid: false},
		{Value: "Груша", Params: []interface{}{"6"}, IsValid: true},
		{Value: stringT("string str"), Params: []interface{}{"3"}, IsValid: false},
	}

	testItems(t, Max, items)
}

func TestLen(t *testing.T) {
	var items = []testItem{
		{Value: "ABC", Params: []interface{}{3}, IsValid: true},
		{Value: []int{1, 2}, Params: []interface{}{"2"}, IsValid: true},
		{Value: [2]int{1, 2}, Params: []interface{}{9}, IsValid: false},
	}

	testItems(t, Len, items)
}

func TestIn(t *testing.T) {
	var items = []testItem{
		{Value: "yes", Params: []interface{}{"yes", "no"}, IsValid: true},
		{Value: "nO", Params: []interface{}{"yes", "no"}, IsValid: false},
		{Value: -4, Params: []interface{}{"-9", "-4"}, IsValid: true},
		{Value: "Слива", Params: []interface{}{"слива", "персик", "смородина"}, IsValid: false},
		{Value: stringT("yes"), Params: []interface{}{"yes", "no"}, IsValid: true},
		{Value: 1, Params: []interface{}{1, 2, 3}, IsValid: true},
		{Value: 2.0, Params: []interface{}{1, 2, 3}, IsValid: true},
	}

	testItems(t, In, items)
}

func TestNotIn(t *testing.T) {
	var items = []testItem{
		{Value: "A", Params: []interface{}{"A", "B"}, IsValid: false},
		{Value: 1, Params: []interface{}{1, 2}, IsValid: false},
		{Value: 3, Params: []interface{}{1, 2}, IsValid: true},
		{Value: 3, Params: []interface{}{1, 2}, IsValid: true},
	}

	testItems(t, NotIn, items)
}

func TestHasPrefix(t *testing.T) {
	var items = []testItem{
		{Value: "prepare", Params: []interface{}{"pre"}, IsValid: true},
		{Value: "prepare", Params: []interface{}{"PRE"}, IsValid: false},
		{Value: "prepare", Params: []interface{}{"pres"}, IsValid: false},
	}

	testItems(t, HasPrefix, items)
}

func TestHasSuffix(t *testing.T) {
	var items = []testItem{
		{Value: "prepare", Params: []interface{}{"are"}, IsValid: true},
		{Value: "prepare", Params: []interface{}{"ARE"}, IsValid: false},
		{Value: "prepare", Params: []interface{}{"ares"}, IsValid: false},
	}

	testItems(t, HasSuffix, items)
}

func TestRegex(t *testing.T) {
	var items = []testItem{
		{Value: "123", Params: []interface{}{"^[\\d]{3}$"}, IsValid: true},
		{Value: "YES", Params: []interface{}{"(?i)yes"}, IsValid: true},
		{Value: "YES", Params: []interface{}{"yes"}, IsValid: false},
	}

	testItems(t, Regex, items)
}

func TestAlpha(t *testing.T) {
	var items = []testItem{
		{Value: "abc", IsValid: true},
		{Value: "abcXYz", IsValid: true},
		{Value: "abcАБВ", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "123abc", IsValid: false},
		{Value: "  abc  ", IsValid: false},
		{Value: "a b c", IsValid: false},
		{Value: "123", IsValid: false},
	}

	testItems(t, Alpha, items)
}

func TestAlphaNumeric(t *testing.T) {
	var items = []testItem{
		{Value: "abc", IsValid: true},
		{Value: "abc123", IsValid: true},
		{Value: "123", IsValid: true},
		{Value: "", IsValid: false},
		{Value: "abc234/", IsValid: false},
	}

	testItems(t, AlphaNumeric, items)
}

func TestAlphaUnder(t *testing.T) {
	var items = []testItem{
		{Value: "abc_xyz", IsValid: true},
		{Value: "abc", IsValid: true},
		{Value: "_", IsValid: true},
		{Value: "__abc__", IsValid: true},
		{Value: "", IsValid: false},
		{Value: "abc_234", IsValid: false},
		{Value: "435", IsValid: false},
		{Value: "--abc_", IsValid: false},
	}

	testItems(t, AlphaUnder, items)
}

func TestAlphaDash(t *testing.T) {
	var items = []testItem{
		{Value: "abc-xyz", IsValid: true},
		{Value: "abc", IsValid: true},
		{Value: "-", IsValid: true},
		{Value: "--abc--", IsValid: true},
		{Value: "", IsValid: false},
		{Value: "abc_234", IsValid: false},
		{Value: "435", IsValid: false},
		{Value: "--abc_", IsValid: false},
	}

	testItems(t, AlphaDash, items)
}

func TestASCII(t *testing.T) {
	var items = []testItem{
		{Value: " !\"#$%&\\'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~", IsValid: true},
		{Value: "АБВ", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "호랑이", IsValid: false},
	}

	testItems(t, ASCII, items)
}

func TestInt(t *testing.T) {
	var items = []testItem{
		{Value: "123", IsValid: true},
		{Value: "0", IsValid: true},
		{Value: "0.123", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "  1.5", IsValid: false},
		{Value: "1,5", IsValid: false},
	}

	testItems(t, Int, items)
}

func TestFloat(t *testing.T) {
	var items = []testItem{
		{Value: "0.123", IsValid: true},
		{Value: "123", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "  1.5", IsValid: false},
		{Value: "1,5", IsValid: false},
	}

	testItems(t, Float, items)
}

func TestJSON(t *testing.T) {
	var items = []testItem{
		{Value: "\"string value\"", IsValid: true},
		{Value: "[1,4]", IsValid: true},
		{Value: "[]", IsValid: true},
		{Value: "{}", IsValid: true},
		{Value: "{\"id\":5}", IsValid: true},
		{Value: "{'id':5}", IsValid: false},
		{Value: "string", IsValid: false},
		{Value: "{\"string\"}", IsValid: false},
	}

	testItems(t, JSON, items)
}

func TestIp(t *testing.T) {
	var items = []testItem{
		{Value: "127.0.0.1", IsValid: true},
		{Value: "64:ff9b::255.255.255.255", IsValid: true},
		{Value: "400.0.0.0", IsValid: false},
		{Value: "192.168.0.0/16", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "123.456", IsValid: false},
	}

	testItems(t, Ip, items)
}

func TestIpv4(t *testing.T) {
	var items = []testItem{
		{Value: "216.3.128.12", IsValid: true},
		{Value: "127.0.0.1", IsValid: true},
		{Value: "192.168.0.0", IsValid: true},
		{Value: "64:ff9b::255.255.255.255", IsValid: false},
		{Value: "192.168.0.0/16", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "123.456", IsValid: false},
	}

	testItems(t, Ipv4, items)
}

func TestIpv6(t *testing.T) {
	var items = []testItem{
		{Value: "64:ff9b::255.255.255.255", IsValid: true},
		{Value: "FE80:0000:0000:0000:0202:B3FF:FE1E:8329", IsValid: true},
		{Value: "[2001:db8:0:1]:80", IsValid: false},
		{Value: "216.3.128.12", IsValid: false},
		{Value: "127.0.0.1", IsValid: false},
		{Value: "192.168.0.0", IsValid: false},
		{Value: "192.168.0.0/16", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "123.456", IsValid: false},
	}

	testItems(t, Ipv6, items)
}

func TestContains(t *testing.T) {
	var items = []testItem{
		{Value: "Prepare", Params: []interface{}{"rep"}, IsValid: true},
		{Value: "Prepare", Params: []interface{}{"Pre"}, IsValid: true},
		{Value: "Prepare", Params: []interface{}{"are"}, IsValid: true},
		{Value: "Prepare", Params: []interface{}{""}, IsValid: true},
		{Value: "", Params: []interface{}{""}, IsValid: true},
		{Value: "Prepare", Params: []interface{}{"res"}, IsValid: false},
		{Value: "Prepare", Params: []interface{}{"aRe"}, IsValid: false},
		{Value: "Prepare", Params: []interface{}{"pre"}, IsValid: false},
	}

	testItems(t, Contains, items)
}

func TestGt(t *testing.T) {
	var items = []testItem{
		{Value: "ABC", Params: []interface{}{2}, IsValid: true},
		{Value: 10, Params: []interface{}{2}, IsValid: true},
		{Value: []int{1, 2, 3}, Params: []interface{}{"2"}, IsValid: true},
		{Value: "ABC", Params: []interface{}{3}, IsValid: false},
		{Value: "ABC", Params: []interface{}{5}, IsValid: false},
		{Value: "", Params: []interface{}{2}, IsValid: false},
		{Value: 5, Params: []interface{}{20}, IsValid: false},
	}

	testItems(t, Gt, items)
}

func TestLt(t *testing.T) {
	var items = []testItem{
		{Value: "ABC", Params: []interface{}{4}, IsValid: true},
		{Value: 10, Params: []interface{}{40}, IsValid: true},
		{Value: []int{1, 2, 3}, Params: []interface{}{"4"}, IsValid: true},
		{Value: "", Params: []interface{}{2}, IsValid: true},
		{Value: "ABC", Params: []interface{}{3}, IsValid: false},
		{Value: "ABC", Params: []interface{}{1}, IsValid: false},
		{Value: 50, Params: []interface{}{20}, IsValid: false},
	}

	testItems(t, Lt, items)
}

func TestHasKeys(t *testing.T) {

	m := map[string]bool{
		"x": true,
		"y": false,
		"z": true,
	}

	var items = []testItem{
		{Value: m, Params: []interface{}{"x", "y"}, IsValid: true},
		{Value: m, Params: []interface{}{"x", "y", "f"}, IsValid: false},
		{Value: m, Params: []interface{}{""}, IsValid: false},
	}

	testItems(t, HasKeys, items)
}

func TestHasOnlyKeys(t *testing.T) {

	m := map[string]bool{
		"x": true,
		"y": false,
		"z": true,
	}

	var items = []testItem{
		{Value: m, Params: []interface{}{"x", "y", "z"}, IsValid: true},
		{Value: m, Params: []interface{}{"x", "y"}, IsValid: false},
		{Value: m, Params: []interface{}{"x", "y", "f"}, IsValid: false},
		{Value: m, Params: []interface{}{""}, IsValid: false},
	}

	testItems(t, HasOnlyKeys, items)
}

func TestTime(t *testing.T) {
	var items = []testItem{
		{Value: "12:30:01", IsValid: true},
		{Value: "00:00:00", IsValid: true},
		{Value: "00:00:90", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "abc", IsValid: false},
		{Value: "12:00", IsValid: false},
	}

	testItems(t, Time, items)
}

func TestUpperCase(t *testing.T) {
	var items = []testItem{
		{Value: "ABC", IsValid: true},
		{Value: "", IsValid: true},
		{Value: "АБВ", IsValid: true},
		{Value: " ", IsValid: true},
		{Value: "13 - 54.5", IsValid: true},
		{Value: "ABC ", IsValid: true},
		{Value: "abc", IsValid: false},
		{Value: "abcABC", IsValid: false},
	}

	testItems(t, UpperCase, items)
}

func TestLowerCase(t *testing.T) {
	var items = []testItem{
		{Value: "abc", IsValid: true},
		{Value: "", IsValid: true},
		{Value: "абв", IsValid: true},
		{Value: " ", IsValid: true},
		{Value: "13 - 54.5", IsValid: true},
		{Value: "abc ", IsValid: true},
		{Value: "ABC", IsValid: false},
		{Value: "abcABC", IsValid: false},
	}

	testItems(t, LowerCase, items)
}

func TestPassword(t *testing.T) {
	var items = []testItem{
		{Value: "abcABC0123", IsValid: true},
		{Value: "abcABC0123#24!!!", IsValid: true},
		{Value: "abcA01", IsValid: false},
		{Value: "abczyxfsdfjsf", IsValid: false},
		{Value: "abczyxfsdfjsf123123", IsValid: false},
		{Value: "SDFSFKLSFLSF123123", IsValid: false},
	}

	testItems(t, Password, items)
}

func TestDate(t *testing.T) {
	var items = []testItem{
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006"}, IsValid: true},
		{Value: "01-52-2019", Params: []interface{}{"02-01-2006"}, IsValid: false},
		{Value: "fake str", Params: []interface{}{"02-01-2006"}, IsValid: false},
	}

	testItems(t, Date, items)
}

func TestDateGte(t *testing.T) {
	now := time.Now()

	var items = []testItem{
		{Value: now.Format("02-01-2006"), Params: []interface{}{"02-01-2006", "-1D"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "01-11-2017"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "01-12-2019"}, IsValid: true},
		{Value: now.Format("02-01-2006"), Params: []interface{}{"02-01-2006", "+1D"}, IsValid: false},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
		{Value: "01-52-2019", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
		{Value: "fake str", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
	}

	testItems(t, DateGte, items)
}

func TestDateLte(t *testing.T) {
	now := time.Now()

	var items = []testItem{
		{Value: now.Format("02-01-2006"), Params: []interface{}{"02-01-2006", "+1D"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "10-12-2019"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "01-12-2019"}, IsValid: true},
		{Value: now.Format("02-01-2006"), Params: []interface{}{"02-01-2006", "-1D"}, IsValid: false},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "02-11-2019"}, IsValid: false},
		{Value: "01-52-2019", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
		{Value: "fake str", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
	}

	testItems(t, DateLte, items)
}

func TestDateGt(t *testing.T) {
	now := time.Now()

	var items = []testItem{
		{Value: now.Format("02-01-2006"), Params: []interface{}{"02-01-2006", "-1D"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "10-12-2010"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "01-12-2019"}, IsValid: false},
		{Value: "01-52-2019", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
		{Value: "fake str", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
	}

	testItems(t, DateGt, items)
}

func TestDateLt(t *testing.T) {
	now := time.Now()

	var items = []testItem{
		{Value: now.Format("02-01-2006"), Params: []interface{}{"02-01-2006", "+1D"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "10-12-2020"}, IsValid: true},
		{Value: "01-12-2019", Params: []interface{}{"02-01-2006", "01-12-2019"}, IsValid: false},
		{Value: "01-52-2019", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
		{Value: "fake str", Params: []interface{}{"02-01-2006", "02-12-2019"}, IsValid: false},
	}

	testItems(t, DateLt, items)
}

func TestCountryCode2(t *testing.T) {
	var items = []testItem{
		{Value: "RU", IsValid: true},
		{Value: "ZU", IsValid: false},
		{Value: "ru", IsValid: false},
		{Value: "RUS", IsValid: false},
	}

	testItems(t, CountryCode2, items)
}

func TestCountryCode3(t *testing.T) {
	var items = []testItem{
		{Value: "AFG", IsValid: true},
		{Value: "RU", IsValid: false},
		{Value: "ZU", IsValid: false},
		{Value: "rus", IsValid: false},
		{Value: "ZUZ", IsValid: false},
	}

	testItems(t, CountryCode3, items)
}

func TestCurrencyCode(t *testing.T) {
	var items = []testItem{
		{Value: "RUB", IsValid: true},
		{Value: "rub", IsValid: false},
		{Value: "RU", IsValid: false},
	}

	testItems(t, CurrencyCode, items)
}

func TestLanguageCode2(t *testing.T) {
	var items = []testItem{
		{Value: "ru", IsValid: true},
		{Value: "zu", IsValid: true},
		{Value: "RU", IsValid: false},
		{Value: "rus", IsValid: false},
		{Value: "ZUZ", IsValid: false},
	}

	testItems(t, LanguageCode2, items)
}

func TestLanguageCode3(t *testing.T) {
	var items = []testItem{
		{Value: "rus", IsValid: true},
		{Value: "RU", IsValid: false},
		{Value: "ru", IsValid: false},
	}

	testItems(t, LanguageCode3, items)
}

func TestCreditCard(t *testing.T) {
	var items = []testItem{
		{Value: "4111111111111111", IsValid: true},            // Visa
		{Value: "5500000000000004", IsValid: true},            // MasterCard
		{Value: "340000000000009", IsValid: true},             // American Express
		{Value: "30000000000004", IsValid: true},              // Diner's Club
		{Value: "6011000000000004", IsValid: true},            // Discover
		{Value: "201400000000009", IsValid: true},             // en Route
		{Value: "3088000000000009", IsValid: true},            // JCB
		{Value: "     4111111111111111     ", IsValid: false}, // Visa
		{Value: "0000 1111 1111 1111", IsValid: false},
		{Value: "4111 1111 1111 1111", IsValid: false},
		{Value: "4111 1111 1111", IsValid: false},
		{Value: "", IsValid: false},
		{Value: "4111 1111 1111 1111 1111 1111 1111 1111", IsValid: false},
	}

	testItems(t, CreditCard, items)
}

func testItems(t *testing.T, fn Validator, items []testItem) {
	for _, item := range items {
		err := fn(reflect.ValueOf(item.Value), item.Options, item.Params...)
		if (err != nil && item.IsValid) || (err == nil && item.IsValid == false) {
			t.Error(fmt.Sprintf("on value «%s» params: «%s»", item.Value, item.Params))
		}
	}
}
