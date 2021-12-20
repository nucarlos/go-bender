package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var ISO4217List = []string{
	"AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN",
	"BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL", "BSD", "BTN", "BWP", "BYN", "BZD",
	"CAD", "CDF", "CHE", "CHF", "CHW", "CLF", "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK",
	"DJF", "DKK", "DOP", "DZD",
	"EGP", "ERN", "ETB", "EUR",
	"FJD", "FKP",
	"GBP", "GEL", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD",
	"HKD", "HNL", "HRK", "HTG", "HUF",
	"IDR", "ILS", "INR", "IQD", "IRR", "ISK",
	"JMD", "JOD", "JPY",
	"KES", "KGS", "KHR", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT",
	"LAK", "LBP", "LKR", "LRD", "LSL", "LYD",
	"MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRO", "MUR", "MVR", "MWK", "MXN", "MXV", "MYR", "MZN",
	"NAD", "NGN", "NIO", "NOK", "NPR", "NZD",
	"OMR",
	"PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PYG",
	"QAR",
	"RON", "RSD", "RUB", "RWF",
	"SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SOS", "SRD", "SSP", "STD", "STN", "SVC", "SYP", "SZL",
	"THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TWD", "TZS",
	"UAH", "UGX", "USD", "USN", "UYI", "UYU", "UYW", "UZS",
	"VEF", "VES", "VND", "VUV", "VEB",
	"WST",
	"XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XTS", "XUA", "XXX",
	"YER",
	"ZAR", "ZMW", "ZWL",
}

// IsISO4217 checks if string is valid ISO currency code
func ISO4217(str string) bool {
	for _, currency := range ISO4217List {
		if str == currency {
			return true
		}
	}

	return false
}

func MyIsISO4217(value interface{}) error {
	s, _ := value.(string)
	if ISO4217(s) {
		return nil
	}

	return validation.NewError("validation_is_currency_code", "must be valid ISO 4217 currency code")
}
