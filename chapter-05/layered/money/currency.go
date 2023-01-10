package money

import "strings"

// Currency defines the code of a money and its decimal precision.
type Currency struct {
	code      string
	precision int
	toEuro    float32
}

func (c Currency) Code() string {
	return c.code
}

// errUnknownCurrency is returned when a currency is unsupported.
const errUnknownCurrency = moneyError("unknown currency")

// currencies defines the supported currencies.
var currencies = map[string]Currency{
	"EUR": {code: "EUR", precision: 2},
	"USD": {code: "USD", precision: 2},
	"JPY": {code: "JPY", precision: 0}, // there is no "cent" for yen
	"BGN": {code: "BGN", precision: 2},
	"CZK": {code: "CZK", precision: 2},
	"DKK": {code: "DKK", precision: 2},
	"GBP": {code: "GBP", precision: 2},
	"HUF": {code: "HUF", precision: 2},
	"PLN": {code: "PLN", precision: 2},
	"RON": {code: "RON", precision: 2},
	"SEK": {code: "SEK", precision: 2},
	"CHF": {code: "CHF", precision: 2},
	"ISK": {code: "ISK", precision: 2},
	"NOK": {code: "NOK", precision: 2},
	"HRK": {code: "HRK", precision: 2},
	"TRY": {code: "TRY", precision: 2},
	"AUD": {code: "AUD", precision: 2},
	"BRL": {code: "BRL", precision: 2},
	"CAD": {code: "CAD", precision: 2},
	"CNY": {code: "CNY", precision: 2},
	"HKD": {code: "HKD", precision: 2},
	"IDR": {code: "IDR", precision: 2},
	"ILS": {code: "ILS", precision: 2},
	"INR": {code: "INR", precision: 2},
	"KRW": {code: "KRW", precision: 2},
	"MXN": {code: "MXN", precision: 2},
	"MYR": {code: "MYR", precision: 2},
	"NZD": {code: "NZD", precision: 2},
	"PHP": {code: "PHP", precision: 2},
	"SGD": {code: "SGD", precision: 2},
	"THB": {code: "THB", precision: 2},
	"ZAR": {code: "ZAR", precision: 2},
}

// parseCurrency returns the currency associated to a name and may return errUnknownCurrency.
func parseCurrency(code string) (Currency, error) {
	// Make sure we are case agnostic by transforming the input to uppercase.
	c, ok := currencies[strings.ToUpper(code)]
	if !ok {
		return Currency{}, errUnknownCurrency
	}
	return c, nil
}
