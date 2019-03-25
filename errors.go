package validation

import "encoding/json"

// The value's validation errors list
type ErrorList []error

// The struct's validation errors map
type ErrorMap map[string]ErrorList

// Checks that the errors list is empty
func (e ErrorList) Empty() bool {
	return len(e) == 0
}

// Marshaling for ErrorList
func (e ErrorList) MarshalJSON() ([]byte, error) {
	res := make([]string, len(e))
	for i, v := range e {
		res[i] = v.Error()
	}
	return json.Marshal(res)
}

// Returns JSON representation of ErrorList
func (e ErrorList) JSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// Checks that the errors map is empty
func (e ErrorMap) Empty() bool {
	for _, item := range e {
		if !item.Empty() {
			return false
		}
	}
	return true
}

// Returns JSON representation of ErrorMap
func (e ErrorMap) JSON() string {
	b, _ := json.Marshal(e)
	return string(b)
}
