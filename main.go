package validation

import (
	"errors"
	"fmt"
	"reflect"
)

// Wrapper for pass custom validators in functional way
type Wrapper struct {
	Function  Validator
	Params    []interface{}
	Reflected bool
}

var defaultTag = "valid"

// Validate scalar value
func ValidateValue(value interface{}, args ...interface{}) (errors []error) {
	wrappers, options, actions := prepareRules(args...)

	for _, action := range actions {
		value = action(value)
	}

	reflectedValue := valueOf(value)

	if options.Has(Ignore) {
		return errors
	}

	if options.Has(Required) == false {
		if Empty(reflectedValue, OptionList{}) == nil {
			return errors
		}
	}

	for _, wrapper := range wrappers {
		var err error
		if wrapper.Reflected {
			err = wrapper.Function(reflectedValue, options, wrapper.Params)
		} else {
			err = wrapper.Function(value, options, wrapper.Params)
		}
		if err != nil {
			errors = append(errors, err)
			if options.Has(Lazy) {
				return errors
			}
		}
	}

	return errors
}

func ValidateStruct(s interface{}, tags ...string) map[string][]error {
	errs := make(map[string][]error)
	fields, err := InspectStruct(s, tags...)
	if err != nil {
		panic("struct only") //todo better solution
	}

	for _, field := range fields {
		fieldErrs := ValidateValue(field.Value, field.ValidTag)
		if fieldErrs != nil {
			errs[field.Name] = fieldErrs
		}
	}

	return errs
}

func InspectStruct(s interface{}, tags ...string) (res []StructField, err error) {
	typeOf := reflect.TypeOf(s)
	valueOf := reflect.ValueOf(s)

	if typeOf.Kind() != reflect.Struct {
		return nil, errors.New("value must be a struct")
	}

	for i := 0; i < typeOf.NumField(); i++ {
		f := StructField{
			Name:  typeOf.Field(i).Name,
			Value: valueOf.Field(i),
		}

		if len(tags) > 0 {
			for _, tag := range tags {
				tagValue := typeOf.Field(i).Tag.Get(tag)
				if len(tagValue) > 0 && len(f.ValidTag) > 0 {
					f.ValidTag = f.ValidTag + "|" + tagValue
				} else {
					f.ValidTag = tagValue
				}
			}
		} else {
			f.ValidTag = typeOf.Field(i).Tag.Get(defaultTag)
		}

		res = append(res, f)
	}

	return res, nil
}

func By(function Validator, params ...interface{}) Wrapper {
	return Wrapper{Function: function, Params: params}
}

// Find and grouping rules by validators, options, actions
func prepareRules(args ...interface{}) ([]Wrapper, OptionList, ActionMap) {
	var wrappers []Wrapper
	var options OptionList
	var actions = make(ActionMap)

	prepareRule := func(rule Rule, wrp *[]Wrapper, options *OptionList, actions ActionMap) {
		if validator, ok := rule.Validator(); ok {
			*wrp = append(*wrp, Wrapper{Function: validator, Params: rule.Params, Reflected: rule.IsBuiltin()})
		}
		if option, ok := rule.Option(); ok {
			*options = append(*options, option)
		}
		if action, ok := rule.Action(); ok {
			actions[actions.key()] = action
		}
		panic(fmt.Sprintf("Rule \"%s\" not found", rule.Name))
	}

	for _, arg := range args {
		switch arg.(type) {

		case string:
			rules := Parse(arg.(string))
			for _, rule := range rules {
				prepareRule(rule, &wrappers, &options, actions)
			}

		case Validator:
			wrappers = append(wrappers, Wrapper{Function: arg.(Validator)})

		case Wrapper:
			wrappers = append(wrappers, arg.(Wrapper))

		case Option:
			options = append(options, arg.(Option))

		case Action:
			actions.Add(actions.key(), arg.(Action))

		case Rule:
			prepareRule(arg.(Rule), &wrappers, &options, actions)

		case func(interface{}, OptionList, ...interface{}) error:
			function := arg.(func(interface{}, OptionList, ...interface{}) error)
			wrappers = append(wrappers, Wrapper{Function: function})

		case func(interface{}) interface{}:
			action := arg.(func(interface{}) interface{})
			actions[actions.key()] = action

		default:
			panic(errorWrongType + fmt.Sprintf("%T", arg))
		}
	}

	return wrappers, options, actions
}
