package validation

import (
	"errors"
	"testing"
)

func TestErrorList_Empty(t *testing.T) {
	tru := ErrorList{}
	fal := ErrorList{errors.New("not empty")}

	if tru.Empty() == false || fal.Empty() == true {
		t.Fail()
	}
}

func TestErrorList_JSON(t *testing.T) {
	r := "[\"error\"]"
	e := ErrorList{errors.New("error")}

	if e.JSON() != r {
		t.Fail()
	}
}

func TestErrorMap_Empty(t *testing.T) {
	tru := ErrorMap{
		"key": ErrorList{},
	}
	fal := ErrorMap{
		"key": ErrorList{errors.New("error")},
	}

	if tru.Empty() == false || fal.Empty() == true {
		t.Fail()
	}
}

func TestErrorMap_JSON(t *testing.T) {
	r := "{\"key\":[\"error\"]}"
	e := ErrorMap{
		"key": ErrorList{errors.New("error")},
	}

	if e.JSON() != r {
		t.Fail()
	}
}
