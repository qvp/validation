package validation

// Map of validation error messages
type messages map[string]string

var Messages = messages{
	"required":      "is required",
	"min":           "must be greater or equal of {0}",
	"max":           "must be lower or equal of {0}",
	"in":            "must be in {0}",
	"email":         "must be a valid email address",
	"url":           "must be a valid url",
	"date":          "must be a valid date in {0} format",
	"country_code2": "must be a valid country code in AA format",
	"country_code3": "must be a valid country code in AAA format",
}

// Add validation error message
func (m messages) Add(rule string, message string) {
	m[rule] = message
}

// Add or replace validation error messages
func (m messages) Update(messages map[string]string) {
	for rule, message := range messages {
		m[rule] = message
	}
}
