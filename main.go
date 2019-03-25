package validation

import (
	"fmt"
	"reflect"
)

// Wrapper for pass custom validators in functional way
type Wrapper struct {
	Function  Validator
	Params    []interface{}
	Reflected bool
}

// Represent struct attribute for validation
type Field struct {
	Name  string
	Value reflect.Value
	Rules string
}

var defaultTag = "valid"

// Validate structure
func ValidateStruct(s interface{}, tags ...string) ErrorMap {
	errs := ErrorMap{}
	fields := InspectStruct(s, tags...)

	for _, field := range fields {
		fieldErrs := validate(s, field.Value, field.Rules)
		if fieldErrs != nil {
			errs[field.Name] = fieldErrs
		}
	}

	return errs
}

// Validate scalar value
func ValidateValue(value interface{}, args ...interface{}) ErrorList {
	return validate(value, value, args...)
}

func InspectStruct(s interface{}, tags ...string) (res []Field) {
	typeOf := reflect.TypeOf(s)
	valueOf := reflect.ValueOf(s)

	if typeOf.Kind() != reflect.Struct {
		panic(errorWrongType)
	}

	for i := 0; i < typeOf.NumField(); i++ {
		f := Field{
			Name:  typeOf.Field(i).Name,
			Value: valueOf.Field(i),
		}

		if len(tags) > 0 {
			for _, tag := range tags {
				tagValue := typeOf.Field(i).Tag.Get(tag)
				if len(tagValue) > 0 && len(f.Rules) > 0 {
					f.Rules = f.Rules + "|" + tagValue
				} else {
					f.Rules = tagValue
				}
			}
		} else {
			f.Rules = typeOf.Field(i).Tag.Get(defaultTag)
		}

		res = append(res, f)
	}

	return res
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

		} else if option, ok := rule.Option(); ok {
			*options = append(*options, option)

		} else if action, ok := rule.Action(); ok {
			actions[actions.key()] = action

		} else {
			panic(fmt.Sprintf("Rule \"%s\" not found", rule.Name))
		}
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
			fmt.Printf("###### %T", arg)
			panic(errorWrongType)
		}
	}

	return wrappers, options, actions
}

func validate(fullValue interface{}, value interface{}, args ...interface{}) ErrorList {
	var errs ErrorList
	wrappers, options, actions := prepareRules(args...)

	for _, action := range actions {
		value = action(value)
	}

	reflectedValue := valueOf(value)

	if options.Has(Ignore) {
		return ErrorList{}
	}

	if options.Has(Required) == false {
		if empty(reflectedValue, OptionList{}) == nil {
			return ErrorList{}
		}
	}

	for _, wrapper := range wrappers {
		var err error
		if wrapper.Reflected {
			err = wrapper.Function(reflectedValue, options, wrapper.Params)
		} else {
			err = wrapper.Function(fullValue, options, wrapper.Params)
		}
		if err != nil {
			errs = append(errs, err)
			if options.Has(Lazy) {
				return errs
			}
		}
	}

	return errs
}
