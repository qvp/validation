package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	Now       DatePlaceholder = "now"
	Today     DatePlaceholder = "today"
	Yesterday DatePlaceholder = "yesterday"
	Tomorrow  DatePlaceholder = "tomorrow"
	PreviousWeek DatePlaceholder = "previous_week"
	PreviousMonth DatePlaceholder = "previous_month"
	PreviousYear DatePlaceholder = "previous_year"
	CurrentWeek  DatePlaceholder = "current_week"
	CurrentMonth DatePlaceholder = "current_month"
	CurrentYear  DatePlaceholder = "current_year"
	NextWeek DatePlaceholder

	YearsAgo18 DatePlaceholder = "18_years_ago"
	YearsAgo21 DatePlaceholder = "21_years_ago"
)

// Return length or value converted to float
func size(value reflect.Value) float64 {
	switch value.Kind() {
	case reflect.String:
		return float64(len([]rune(value.String())))
	case reflect.Array, reflect.Map, reflect.Chan, reflect.Slice:
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

// Return date by placeholder
func GetDate(placeholder DatePlaceholder) (time.Time, error) {
	now := time.Now()

	switch placeholder {
	case Now:
		return now, nil
	case Today:
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0,0, 0, now.Location()), nil
	case Yesterday:
		return now.AddDate(0, 0, -1), nil
	case Tomorrow:
		return now.AddDate(0, 0, 1), nil
	case YearsAgo18:
		return now.AddDate(-18, 0, 0), nil
	case YearsAgo21:
		return now.AddDate(-21, 0, 0), nil
	case ThisWeek:
		day := now.Day() - int(now.Weekday())
		return time.Date(now.Year(), now.Month(), day, 0, 0,0, 0, now.Location()), nil
	case ThisMonth:
		return time.Date(now.Year(), now.Month(), 1, 0, 0,0, 0, now.Location()), nil
	case ThisYear:
		return time.Date(now.Year(), 1, 1, 0, 0,0, 0, now.Location()), nil
	default:
		return time.Time{}, errors.New("placeholder not exists")
	}
}
