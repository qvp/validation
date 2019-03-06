package validation

var (
	CountryCodes2 = []string{
		"AF", "AX", "AL", "DZ", "AS", "AD", "AO", "AI", "AQ", "AG", "AR", "AM", "AW", "AU", "AT", "AZ", "BS", "BH",
		"BD", "BB", "BY", "BE", "BZ", "BJ", "BM", "BT", "BO", "BQ", "BA", "BW", "BV", "BR", "IO", "BN", "BG", "BF",
		"BI", "CV", "KH", "CM", "CA", "KY", "CF", "TD", "CL", "CN", "CX", "CC", "CO", "KM", "CG", "CD", "CK", "CR",
		"CI", "HR", "CU", "CW", "CY", "CZ", "DK", "DJ", "DM", "DO", "EC", "EG", "SV", "GQ", "ER", "EE", "SZ", "ET",
		"FK", "FO", "FJ", "FI", "FR", "GF", "PF", "TF", "GA", "GM", "GE", "DE", "GH", "GI", "GR", "GL", "GD", "GP",
		"GU", "GT", "GG", "GN", "GW", "GY", "HT", "HM", "VA", "HN", "HK", "HU", "IS", "IN", "ID", "IR", "IQ", "IE",
		"IM", "IL", "IT", "JM", "JP", "JE", "JO", "KZ", "KE", "KI", "KP", "KR", "KW", "KG", "LA", "LV", "LB", "LS",
		"LR", "LY", "LI", "LT", "LU", "MO", "MK", "MG", "MW", "MY", "MV", "ML", "MT", "MH", "MQ", "MR", "MU", "YT",
		"MX", "FM", "MD", "MC", "MN", "ME", "MS", "MA", "MZ", "MM", "NA", "NR", "NP", "NL", "NC", "NZ", "NI", "NE",
		"NG", "NU", "NF", "MP", "NO", "OM", "PK", "PW", "PS", "PA", "PG", "PY", "PE", "PH", "PN", "PL", "PT", "PR",
		"QA", "RE", "RO", "RU", "RW", "BL", "SH", "KN", "LC", "MF", "PM", "VC", "WS", "SM", "ST", "SA", "SN", "RS",
		"SC", "SL", "SG", "SX", "SK", "SI", "SB", "SO", "ZA", "GS", "SS", "ES", "LK", "SD", "SR", "SJ", "SE", "CH",
		"SY", "TW", "TJ", "TZ", "TH", "TL", "TG", "TK", "TO", "TT", "TN", "TR", "TM", "TC", "TV", "UG", "UA", "AE",
		"GB", "US", "UM", "UY", "UZ", "VU", "VE", "VN", "VG", "VI", "WF", "EH", "YE", "ZM", "ZW",
	}

	CountryCodes3 = []string{
		"AFG", "ALA", "ALB", "DZA", "ASM", "AND", "AGO", "AIA", "ATA", "ATG", "ARG", "ARM", "ABW", "AUS", "AUT",
		"AZE", "BHS", "BHR", "BGD", "BRB", "BLR", "BEL", "BLZ", "BEN", "BMU", "BTN", "BOL", "BES", "BIH", "BWA",
		"BVT", "BRA", "IOT", "BRN", "BGR", "BFA", "BDI", "CPV", "KHM", "CMR", "CAN", "CYM", "CAF", "TCD", "CHL",
		"CHN", "CXR", "CCK", "COL", "COM", "COG", "COD", "COK", "CRI", "CIV", "HRV", "CUB", "CUW", "CYP", "CZE",
		"DNK", "DJI", "DMA", "DOM", "ECU", "EGY", "SLV", "GNQ", "ERI", "EST", "SWZ", "ETH", "FLK", "FRO", "FJI",
		"FIN", "FRA", "GUF", "PYF", "ATF", "GAB", "GMB", "GEO", "DEU", "GHA", "GIB", "GRC", "GRL", "GRD", "GLP",
		"GUM", "GTM", "GGY", "GIN", "GNB", "GUY", "HTI", "HMD", "VAT", "HND", "HKG", "HUN", "ISL", "IND", "IDN",
		"IRN", "IRQ", "IRL", "IMN", "ISR", "ITA", "JAM", "JPN", "JEY", "JOR", "KAZ", "KEN", "KIR", "PRK", "KOR",
		"KWT", "KGZ", "LAO", "LVA", "LBN", "LSO", "LBR", "LBY", "LIE", "LTU", "LUX", "MAC", "MKD", "MDG", "MWI",
		"MYS", "MDV", "MLI", "MLT", "MHL", "MTQ", "MRT", "MUS", "MYT", "MEX", "FSM", "MDA", "MCO", "MNG", "MNE",
		"MSR", "MAR", "MOZ", "MMR", "NAM", "NRU", "NPL", "NLD", "NCL", "NZL", "NIC", "NER", "NGA", "NIU", "NFK",
		"MNP", "NOR", "OMN", "PAK", "PLW", "PSE", "PAN", "PNG", "PRY", "PER", "PHL", "PCN", "POL", "PRT", "PRI",
		"QAT", "REU", "ROU", "RUS", "RWA", "BLM", "SHN", "KNA", "LCA", "MAF", "SPM", "VCT", "WSM", "SMR", "STP",
		"SAU", "SEN", "SRB", "SYC", "SLE", "SGP", "SXM", "SVK", "SVN", "SLB", "SOM", "ZAF", "SGS", "SSD", "ESP",
		"LKA", "SDN", "SUR", "SJM", "SWE", "CHE", "SYR", "TWN", "TJK", "TZA", "THA", "TLS", "TGO", "TKL", "TON",
		"TTO", "TUN", "TUR", "TKM", "TCA", "TUV", "UGA", "UKR", "ARE", "GBR", "USA", "UMI", "URY", "UZB", "VUT",
		"VEN", "VNM", "VGB", "VIR", "WLF", "ESH", "YEM", "ZMB", "ZWE",
	}

	CurrencyCodes = []string{
		"AFN", "ALL", "DZD", "USD", "EUR", "AOA", "XCD", "XCD", "ARS", "AMD", "AWG", "AUD", "EUR", "AZN", "BSD",
		"BHD", "BDT", "BBD", "BYR", "EUR", "BZD", "XOF", "BMD", "BTN", "INR", "BOB", "BOV", "USD", "BAM", "BWP",
		"NOK", "BRL", "USD", "BND", "BGN", "XOF", "BIF", "CVE", "KHR", "XAF", "CAD", "KYD", "XAF", "XAF", "CLF",
		"CLP", "CNY", "AUD", "AUD", "COP", "COU", "KMF", "CDF", "XAF", "NZD", "CRC", "HRK", "CUC", "CUP", "ANG",
		"EUR", "CZK", "XOF", "DKK", "DJF", "XCD", "DOP", "USD", "EGP", "SVC", "USD", "XAF", "ERN", "EUR", "ETB",
		"EUR", "FKP", "DKK", "FJD", "EUR", "EUR", "EUR", "XPF", "EUR", "XAF", "GMD", "GEL", "EUR", "GHS", "GIP",
		"EUR", "DKK", "XCD", "EUR", "USD", "GTQ", "GBP", "GNF", "XOF", "GYD", "HTG", "USD", "AUD", "EUR", "HNL",
		"HKD", "HUF", "ISK", "INR", "IDR", "XDR", "IRR", "IQD", "EUR", "GBP", "ILS", "EUR", "JMD", "JPY", "GBP",
		"JOD", "KZT", "KES", "AUD", "KPW", "KRW", "KWD", "KGS", "LAK", "EUR", "LBP", "LSL", "ZAR", "LRD", "LYD",
		"CHF", "EUR", "EUR", "MOP", "MKD", "MGA", "MWK", "MYR", "MVR", "XOF", "EUR", "USD", "EUR", "MRU", "MUR",
		"EUR", "XUA", "MXN", "MXV", "USD", "MDL", "EUR", "MNT", "EUR", "XCD", "MAD", "MZN", "MMK", "NAD", "ZAR",
		"AUD", "NPR", "EUR", "XPF", "NZD", "NIO", "XOF", "NGN", "NZD", "AUD", "USD", "NOK", "OMR", "PKR", "USD",
		"PAB", "USD", "PGK", "PYG", "PEN", "PHP", "NZD", "PLN", "EUR", "USD", "QAR", "RON", "RUB", "RWF", "EUR",
		"EUR", "SHP", "XCD", "XCD", "EUR", "EUR", "XCD", "WST", "EUR", "STN", "SAR", "XOF", "RSD", "SCR", "SLL",
		"SGD", "ANG", "XSU", "EUR", "EUR", "SBD", "SOS", "ZAR", "SSP", "EUR", "LKR", "SDG", "SRD", "NOK", "SZL",
		"SEK", "CHE", "CHF", "CHW", "SYP", "TWD", "TJS", "TZS", "THB", "USD", "XOF", "NZD", "TOP", "TTD", "TND",
		"TRY", "TMT", "USD", "AUD", "UGX", "UAH", "AED", "GBP", "USD", "USD", "USN", "UYI", "UYU", "UZS", "VUV",
		"VEF", "VND", "USD", "USD", "XPF", "MAD", "YER", "ZMW", "ZWL", "EUR",
	}

	LanguageCodes2 = []string{
		"ab", "aa", "af", "sq", "am", "ar", "an", "hy", "as", "ae", "ay", "az", "ba", "eu", "be", "bn", "bh", "bi",
		"bs", "br", "bg", "my", "ca", "ch", "ce", "zh", "cu", "cv", "kw", "co", "hr", "cs", "da", "dv", "nl", "dz",
		"en", "eo", "et", "fo", "fj", "fi", "fr", "gd", "gl", "ka", "de", "el", "gn", "gu", "ht", "ha", "he", "hz",
		"hi", "ho", "hu", "is", "io", "id", "ia", "ie", "iu", "ik", "ga", "it", "ja", "jv", "kl", "kn", "ks", "kk",
		"km", "ki", "rw", "ky", "kv", "ko", "kj", "ku", "lo", "la", "lv", "li", "ln", "lt", "lb", "mk", "mg", "ms",
		"ml", "mt", "gv", "mi", "mr", "mh", "mo", "mn", "na", "nv", "nd", "nr", "ng", "ne", "se", "no", "nb", "nn",
		"ny", "oc", "or", "om", "os", "pi", "pa", "fa", "pl", "pt", "ps", "qu", "rm", "ro", "rn", "ru", "sm", "sg",
		"sa", "sc", "sr", "sn", "ii", "sd", "si", "sk", "sl", "so", "st", "es", "su", "sw", "ss", "sv", "tl", "ty",
		"tg", "ta", "tt", "te", "th", "bo", "ti", "to", "ts", "tn", "tr", "tk", "tw", "ug", "uk", "ur", "uz", "vi",
		"vo", "wa", "cy", "fy", "wo", "xh", "yi", "yo", "za", "zu",
	}

	LanguageCodes3 = []string{
		"abk", "aar", "afr", "alb", "sqi", "amh", "ara", "arg", "arm", "hye", "asm", "ave", "aym", "aze", "bak",
		"baq", "eus", "bel", "ben", "bih", "bis", "bos", "bre", "bul", "bur", "mya", "cat", "cha", "che", "chi",
		"zho", "chu", "chv", "cor", "cos", "scr", "hrv", "cze", "ces", "dan", "div", "dut", "nld", "dzo", "eng",
		"epo", "est", "fao", "fij", "fin", "fre", "fra", "gla", "glg", "geo", "kat", "ger", "deu", "gre", "ell",
		"grn", "guj", "hat", "hau", "heb", "her", "hin", "hmo", "hun", "ice", "isl", "ido", "ind", "ina", "ile",
		"iku", "ipk", "gle", "ita", "jpn", "jav", "kal", "kan", "kas", "kaz", "khm", "kik", "kin", "kir", "kom",
		"kor", "kua", "kur", "lao", "lat", "lav", "lim", "lin", "lit", "ltz", "mac", "mkd", "mlg", "may", "msa",
		"mal", "mlt", "glv", "mao", "mri", "mar", "mah", "mol", "mon", "nau", "nav", "nde", "nbl", "ndo", "nep",
		"sme", "nor", "nob", "nno", "nya", "oci", "ori", "orm", "oss", "pli", "pan", "per", "fas", "pol", "por",
		"pus", "que", "roh", "rum", "ron", "run", "rus", "smo", "sag", "san", "srd", "scc", "srp", "sna", "iii",
		"snd", "sin", "slo", "slk", "slv", "som", "sot", "spa", "sun", "swa", "ssw", "swe", "tgl", "tah", "tgk",
		"tam", "tat", "tel", "tha", "tib", "bod", "tir", "ton", "tso", "tsn", "tur", "tuk", "twi", "uig", "ukr",
		"urd", "uzb", "vie", "vol", "wln", "wel", "cym", "fry", "wol", "xho", "yid", "yor", "zha", "zul",
	}
)
