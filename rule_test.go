package validation

import "testing"

func TestRule_Exists(t *testing.T) {
	true1 := Rule{Name: "required"}
	true2 := Rule{Name: "max"}
	true3 := Rule{Name: "trim"}

	false1 := Rule{Name: "REQUIRED"}
	false2 := Rule{Name: "fake"}

	if !true1.Exists() || !true2.Exists() || !true3.Exists() || false1.Exists() || false2.Exists() {
		t.Error("Error check exists.")
	}
}

func TestRule_Validator(t *testing.T) {
	r1 := Rule{Name: "email"}
	r2 := Rule{Name: "fake"}

	_, true1 := r1.Validator()
	_, false1 := r2.Validator()

	if !true1 || false1 {
		t.Error("Error get validator.")
	}
}

func TestRule_Option(t *testing.T) {
	r1 := Rule{Name: "lazy"}
	r2 := Rule{Name: "fake"}

	_, true1 := r1.Option()
	_, false1 := r2.Option()

	if !true1 || false1 {
		t.Error("Error get option.")
	}
}

func TestRule_Action(t *testing.T) {
	r1 := Rule{Name: "lower"}
	r2 := Rule{Name: "fake"}

	_, true1 := r1.Action()
	_, false1 := r2.Action()

	if !true1 || false1 {
		t.Error("Error get action.")
	}
}

func TestRule_IsBuiltin(t *testing.T) {
	true1 := Rule{Name: "url"}

	false1 := Rule{Name: "required"} // check builtin for options not supported
	false2 := Rule{Name: "trim"}     // check builtin for actions not supported

	Options.Add("custom_option")
	false3 := Rule{Name: "fake"}
	false4 := Rule{Name: "custom_option"}

	if !true1.IsBuiltin() || false1.IsBuiltin() || false2.IsBuiltin() || false3.IsBuiltin() || false4.IsBuiltin() {
		t.Error("Error check built-in.")
	}
}

func TestParse(t *testing.T) { //todo more cases!
	s := "required|max:255|in:x,y,z"
	r := Parse(s)
	if len(r) != 3 {
		t.Error("Error parsing rules.")
	}
}
