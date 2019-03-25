package is

import "github.com/qvp/validation"

func Required(params ...interface{}) validation.Option {
	return validation.Required
}

func Ignore(params ...interface{}) validation.Option {
	return validation.Ignore
}

func Lazy(params ...interface{}) validation.Option {
	return validation.Lazy
}

func Empty(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "empty", Params: params}
}

func Email(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "email", Params: params}
}

func URL(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "url", Params: params}
}

func Accepted(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "accepted", Params: params}
}

func Alpha(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "alpha", Params: params}
}

func AlphaUnder(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "alpha_under", Params: params}
}

func Alphadash(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "alpha_dash", Params: params}
}

func ASCII(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "ascii", Params: params}
}

func Int(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "int", Params: params}
}

func Float(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "float", Params: params}
}

func JSON(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "json", Params: params}
}

func Ip(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "ip", Params: params}
}

func Ipv4(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "ipv4", Params: params}
}

func Ipv6(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "ipv6", Params: params}
}

func Time(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "time", Params: params}
}

func UpperCase(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "upper_case", Params: params}
}

func LowerCase(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "lower_case", Params: params}
}

func CountryCode2(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "country_code2", Params: params}
}

func CountryCode3(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "country_code3", Params: params}
}

func CurrencyCode(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "currency_code", Params: params}
}

func LanguageCode2(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "language_code2", Params: params}
}

func LanguageCode3(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "language_code3", Params: params}
}

func CreditCard(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "credit_card", Params: params}
}

func Password(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "password", Params: params}
}

func Min(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "min", Params: params}
}

func Max(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "max", Params: params}
}

func Len(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "len", Params: params}
}

func In(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "in", Params: params}
}

func NotIn(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "not_in", Params: params}
}

func Date(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "date", Params: params}
}

func Regex(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "regex", Params: params}
}

func Contains(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "contains", Params: params}
}

func Gt(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "gt", Params: params}
}

func Lt(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "lt", Params: params}
}

func DateGte(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "date_gte", Params: params}
}

func DateLte(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "date_lte", Params: params}
}

func DateGt(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "date_gt", Params: params}
}

func DateLt(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "date_lt", Params: params}
}

func HasPrefix(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "has_prefix", Params: params}
}

func HasSuffix(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "has_suffix", Params: params}
}

func HasKeys(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "has_keys", Params: params}
}

func HasOnlyKeys(params ...interface{}) validation.Rule {
	return validation.Rule{Name: "has_only_keys", Params: params}
}
