package validation

import "strings"

// Validation rule
type Rule struct {
	Name   string
	Params []interface{}
}

// Check what rule exists
// Rule can be an Validator or Option or Action
func (r *Rule) Exists() bool {
	if _, exists := r.Validator(); exists {
		return true
	}

	if _, exists := r.Option(); exists {
		return true
	}

	if _, exists := r.Action(); exists {
		return true
	}

	return false
}

// Get Validator if it exists
func (r *Rule) Validator() (Validator, bool) {
	if validator, ok := validators[r.Name]; ok {
		return validator, ok
	}

	if validator, ok := Validators[r.Name]; ok {
		return validator, ok
	}

	return nil, false
}

// Get Option if it exists
func (r *Rule) Option() (Option, bool) {
	for _, o := range Options {
		if r.Name == string(o) {
			return o, true
		}
	}

	return Option(""), false
}

// Get Action if it exists
func (r *Rule) Action() (Action, bool) {
	if action, ok := Actions[r.Name]; ok {
		return action, ok
	}

	return nil, false
}

// Check what validator is built in
// todo maybe Option and Action check to ?
func (r *Rule) IsBuiltin() bool {
	_, ok := validators[r.Name]
	return ok
}

// Parse string of rules
// | - rule separator
// : - split up rule name and parameters
// , - split up parameters
// Example "required|max:255|in:x,y,z"
func Parse(s string) []Rule {
	var r []Rule

	for _, f := range strings.Split(s, "|") {
		p := strings.SplitN(f, ":", 2)
		// todo trim
		if len(p) == 1 {
			r = append(r, Rule{Name: p[0]})
			continue
		}

		if p[0] == "regex" {
			r = append(r, Rule{Name: "regex", Params: []interface{}{p[1]}})
		}

		a := strings.Split(p[1], ",")
		pr := make([]interface{}, len(a))
		for i, v := range a {
			pr[i] = v
		}
		r = append(r, Rule{Name: p[0], Params: pr})
	}
	return r
}
