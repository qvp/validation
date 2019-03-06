package validation

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	validators = map[string]func(reflect.Value) error{
		"empty":          Empty,
		"email":          Email,
		"url":            URL,
		"accepted":       Accepted,
		"alpha":          Alpha,
		"alpha_under":    AlphaUnder,
		"alpha_dash":     AlphaDash,
		"ascii":          ASCII,
		"int":            Int,
		"float":          Float,
		"json":           JSON,
		"ip":             Ip,
		"ipv4":           Ipv4,
		"ipv6":           Ipv6,
		"time":           Time,
		"upper_case":     UpperCase,
		"lower_case":     LowerCase,
		"country_code2":  CountryCode2,
		"country_code3":  CountryCode3,
		"currency_code":  CurrencyCode,
		"language_code2": LanguageCode2,
		"language_code3": LanguageCode3,
		"credit_card":    CreditCard,
		"password":       Password,
	}
	validatorsWithParams = map[string]func(reflect.Value, ...interface{}) error{
		"min":             Min,
		"max":             Max,
		"len":             Len,
		"in":              In,
		"not_in":          NotIn,
		"date":            Date,
		"regex":           Regex,
		"contains":        Contains,
		"gt":              Gt,
		"lt":              Lt,
		"date_gte":        DateGte,
		"date_lte":        DateLte,
		"date_gt":         DateGt,
		"date_lt":         DateLt,
		"has_prefix":      HasPrefix,
		"has_suffix":      HasSuffix,
		"has_keys":        HasKeys,
		"has_only_keys":   HasOnlyKeys,
		"has_values":      HasValues,
		"has_only_values": HasOnlyValues,
	}
	options = map[string]func(reflect.Value) error{
		"required": Required,
		"ignore": Ignore,
		"lazy":   Lazy,
	}
	defaultTag = "valid"
)

func Value(value interface{}, args ...interface{}) (errs []error) {
	refValue, ok := value.(reflect.Value)
	if !ok {
		refValue = reflect.ValueOf(value)
	}

	for _, arg := range args {
		switch arg.(type) {

		case string:
			rules := Parse(arg.(string))
			for _, rule := range rules {
				if function, ok := validators[rule.Name]; ok {
					if err := function(refValue); err != nil {
						errs = append(errs, err)
					}
					continue
				}
				if function, ok := validatorsWithParams[rule.Name]; ok {
					if err := function(refValue, rule.Params...); err != nil {
						errs = append(errs, err)
					}
					continue
				}
				panic(fmt.Sprintf("Validator \"%s\" not found", rule.Name))
			}

		case func(reflect.Value) error:
			function := arg.(func(reflect.Value) error)
			if err := function(refValue); err != nil {
				errs = append(errs, err)
			}

		//case func(interface{}) error: //todo need this type?
		//	function := arg.(func(interface{}) error)
		//	if err := function(refValue); err != nil {
		//		errs = append(errs, err)
		//	}

		case ValidatorWrapper:
			vl := arg.(ValidatorWrapper)
			if err := vl.Function(refValue, vl.Params...); err != nil {
				errs = append(errs, err)
			}

		default:
			panic(fmt.Sprintf("Param must be a string of validator(s) or callable of Validator type. %T", arg))
		}
	}
	return
}

func ValidateStruct(s interface{}, tags ...string) map[string][]error {
	errs := make(map[string][]error)
	fields, err := InspectStruct(s, tags...)
	if err != nil {
		panic("struct only") //todo better solution
	}

	for _, field := range fields {
		fieldErrs := Value(field.Value, field.ValidTag)
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

func By(function ValidatorP, params ...interface{}) ValidatorWrapper {
	return ValidatorWrapper{Function: function, Params: params}
}

func AddValidator(name string, fun interface{}) {
	switch fun.(type) {
	case func(reflect.Value, ...interface{}) error:
		validatorsWithParams[name] = fun.(func(reflect.Value, ...interface{}) error)
		return
	case func(reflect.Value) error:
		validators[name] = fun.(func(reflect.Value) error)
	}

	panic("fun must be of type valid.Validator or valid.ValidatorP")
}

func AddOption() {

}

func AddMessages(messages map[string]string) {

}

func SetDefaultTag(name string) { //todo make this thread safe!!!!
	defaultTag = name
}

func DefaultTag() string {
	return defaultTag
}
