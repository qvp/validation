package validation

import (
	"testing"
)

func TestPrepareRules(t *testing.T) {
	v := func(v interface{}, options OptionList, params ...interface{}) error {
		return nil
	}
	a := func(v interface{}) interface{} {
		return nil
	}

	wrappers, options, actions := prepareRules("required|max:255|lazy|trim|clear", v, Ignore, a)

	if len(wrappers) != 2 {
		t.Error("Error finding validators.")
	}
	if len(options) != 3 {
		t.Error("Error finding options.")
	}
	if len(actions) != 3 {
		t.Error("Error finding actions.")
	}
}
