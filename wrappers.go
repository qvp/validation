package validation

func Empty(params ...interface{}) Rule {
	return Rule{Name: "empty", Params: params}
}

func Email(params ...interface{}) Rule {
	return Rule{Name: "email", Params: params}
}

func URL(params ...interface{}) Rule {
	return Rule{Name: "url", Params: params}
}

func Accepted(params ...interface{}) Rule {
	return Rule{Name: "accepted", Params: params}
}

func Alpha(params ...interface{}) Rule {
	return Rule{Name: "alpha", Params: params}
}

func AlphaUnder(params ...interface{}) Rule {
	return Rule{Name: "alpha_under", Params: params}
}

func Alphadash(params ...interface{}) Rule {
	return Rule{Name: "alpha_dash", Params: params}
}

func ASCII(params ...interface{}) Rule {
	return Rule{Name: "ascii", Params: params}
}

func Int(params ...interface{}) Rule {
	return Rule{Name: "int", Params: params}
}

func Float(params ...interface{}) Rule {
	return Rule{Name: "float", Params: params}
}

func JSON(params ...interface{}) Rule {
	return Rule{Name: "json", Params: params}
}

func Ip(params ...interface{}) Rule {
	return Rule{Name: "ip", Params: params}
}

func Ipv4(params ...interface{}) Rule {
	return Rule{Name: "ipv4", Params: params}
}

func Ipv6(params ...interface{}) Rule {
	return Rule{Name: "ipv6", Params: params}
}

func Time(params ...interface{}) Rule {
	return Rule{Name: "time", Params: params}
}

func UpperCase(params ...interface{}) Rule {
	return Rule{Name: "upper_case", Params: params}
}

func LowerCase(params ...interface{}) Rule {
	return Rule{Name: "lower_case", Params: params}
}

func CountryCode2(params ...interface{}) Rule {
	return Rule{Name: "country_code2", Params: params}
}

func CountryCode3(params ...interface{}) Rule {
	return Rule{Name: "country_code3", Params: params}
}

func CurrencyCode(params ...interface{}) Rule {
	return Rule{Name: "currency_code", Params: params}
}

func LanguageCode2(params ...interface{}) Rule {
	return Rule{Name: "language_code2", Params: params}
}

func LanguageCode3(params ...interface{}) Rule {
	return Rule{Name: "language_code3", Params: params}
}

func CreditCard(params ...interface{}) Rule {
	return Rule{Name: "credit_card", Params: params}
}

func Password(params ...interface{}) Rule {
	return Rule{Name: "password", Params: params}
}

func Min(params ...interface{}) Rule {
	return Rule{Name: "min", Params: params}
}

func Max(params ...interface{}) Rule {
	return Rule{Name: "max", Params: params}
}

func Len(params ...interface{}) Rule {
	return Rule{Name: "len", Params: params}
}

func In(params ...interface{}) Rule {
	return Rule{Name: "in", Params: params}
}

func NotIn(params ...interface{}) Rule {
	return Rule{Name: "not_in", Params: params}
}

func Date(params ...interface{}) Rule {
	return Rule{Name: "date", Params: params}
}

func Regex(params ...interface{}) Rule {
	return Rule{Name: "regex", Params: params}
}

func Contains(params ...interface{}) Rule {
	return Rule{Name: "contains", Params: params}
}

func Gt(params ...interface{}) Rule {
	return Rule{Name: "gt", Params: params}
}

func Lt(params ...interface{}) Rule {
	return Rule{Name: "lt", Params: params}
}

func DateGte(params ...interface{}) Rule {
	return Rule{Name: "date_gte", Params: params}
}

func DateLte(params ...interface{}) Rule {
	return Rule{Name: "date_lte", Params: params}
}

func DateGt(params ...interface{}) Rule {
	return Rule{Name: "date_gt", Params: params}
}

func DateLt(params ...interface{}) Rule {
	return Rule{Name: "date_lt", Params: params}
}

func HasPrefix(params ...interface{}) Rule {
	return Rule{Name: "has_prefix", Params: params}
}

func HasSuffix(params ...interface{}) Rule {
	return Rule{Name: "has_suffix", Params: params}
}

func HasKeys(params ...interface{}) Rule {
	return Rule{Name: "has_keys", Params: params}
}

func HasOnlyKeys(params ...interface{}) Rule {
	return Rule{Name: "has_only_keys", Params: params}
}
