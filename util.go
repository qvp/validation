package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Now       DatePlaceholder = "now"
	Today     DatePlaceholder = "today"
	Yesterday DatePlaceholder = "yesterday"
	Tomorrow  DatePlaceholder = "tomorrow"
)

var (
	regexDateModifier = regexp.MustCompile("([-|+]?[\\d]+)([Y|M|D|h|m|s])")
)

// Return date and time modified by placeholder
// Placeholders list:
// now - current date and current time
// today - current date and 00:00:00:00 time
// yesterday - current date - 1 day and 00:00:00:00 time
// +-{n}Y - add or remove n years to current date and current time
// +-{n}M - add or remove n months to current date and current time
// +-{n}D - add or remove n days to current date and current time
// +-{n}h - add or remove n hours to current date and current time
// +-{n}m - add or remove n minutes to current date and current time
// +-{n}s - add or remove n seconds to current date and current time
//
// You can combine Y,M,D,h,m,s modifiers.
// Example: get a date that was 18 years and 3 days ago, plus 22 seconds:
// dt, err := GetDate("-18Y -3h +22s")
func GetDate(placeholder DatePlaceholder) (time.Time, error) {
	now := time.Now()

	switch placeholder {
	case Now:
		return now, nil
	case Today:
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), nil
	case Yesterday:
		y := now.AddDate(0, 0, -1)
		return time.Date(y.Year(), y.Month(), y.Day(), 0, 0, 0, 0, y.Location()), nil
	case Tomorrow:
		t := now.AddDate(0, 0, 1)
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()), nil
	default:
		modifiers, err := parseDateModifiers(string(placeholder))
		if err == nil {
			date := now.AddDate(modifiers["Y"], modifiers["M"], modifiers["D"])
			h := time.Hour * time.Duration(modifiers["h"])
			m := time.Minute * time.Duration(modifiers["m"])
			s := time.Second * time.Duration(modifiers["s"])
			return date.Add(h + m + s), nil
		}
	}

	return time.Time{}, errors.New("placeholder not exists")
}

// Return length or value converted to float
func size(value reflect.Value) float64 {
	switch value.Kind() {
	case reflect.String:
		return float64(len([]rune(value.String())))
	case reflect.Array, reflect.Map, reflect.Slice:
		return float64(value.Len())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(value.Uint())
	case reflect.Float32, reflect.Float64:
		return value.Float()
	default:
		panic(errorWrongType)
	}
}

// Return new string with replaced parameters by their index
func replace(message string, params []interface{}) string {
	for i, item := range params {
		token := fmt.Sprintf("{%d}", i)
		value := fmt.Sprintf("%v", item)
		message = strings.Replace(message, token, value, -1)
	}
	return message
}

// Return weather that string in slice
func in(value string, items []string) bool {
	for _, item := range items {
		if value == item {
			return true
		}
	}
	return false
}

// Return validation error with replaced parameters
func errorMessage(ruleName string, params ...interface{}) error {
	message := Messages[ruleName]
	if len(message) == 0 {
		message = "validation by " + ruleName + " not pass."
	}
	message = replace(message, params)
	return errors.New(message)
}

// Convert string or digit type to float
func parseFloat(value interface{}) (float64, error) {
	s := parseString(value)
	return strconv.ParseFloat(s, 64)
}

// Convert value to string
func parseString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// Parse strings like a -18Y +3h 22s
// Allowed modifiers: Y, M, D, h, m, s
func parseDateModifiers(s string) (map[string]int, error) {
	var modifiers = map[string]int{"Y": 0, "M": 0, "D": 0, "h": 0, "m": 0, "s": 0}

	matches := regexDateModifier.FindAllStringSubmatch(s, 6)
	if len(matches) == 0 {
		return modifiers, errors.New("modifiers not found")
	}

	for _, item := range matches {
		modifiers[item[2]], _ = strconv.Atoi(item[1])
	}

	return modifiers, nil
}

// Check credit card number by luhn algorithm
func luhn(num string) bool {
	var sum int
	var alternate bool
	var ln = len(num)

	if ln < 13 || ln > 19 {
		return false
	}

	for i := ln - 1; i > -1; i-- {
		n, _ := strconv.Atoi(string(num[i]))
		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}
		alternate = !alternate
		sum += n
	}

	return sum%10 == 0
}
