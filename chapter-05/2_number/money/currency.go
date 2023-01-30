package money

// Currency defines the code of a currency and its decimal precision.
type Currency struct {
	code      string
	precision int
}

// errUnknownCurrency is returned when a currency is unsupported.
const errUnknownCurrency = moneyError("unsupported currency")

// ParseCurrency returns the currency associated to a name and may return errUnknownCurrency.
func ParseCurrency(code string) (Currency, error) {
	switch code {
	case "EUR":
		return Currency{code: "EUR", precision: 2}, nil
	case "USD":
		return Currency{code: "USD", precision: 2}, nil
	case "JPY":
		return Currency{code: "JPY", precision: 0}, nil // there is no "cent" for yen
	case "BGN":
		return Currency{code: "BGN", precision: 2}, nil
	case "CZK":
		return Currency{code: "CZK", precision: 2}, nil
	case "DKK":
		return Currency{code: "DKK", precision: 2}, nil
	case "GBP":
		return Currency{code: "GBP", precision: 2}, nil
	case "HUF":
		return Currency{code: "HUF", precision: 2}, nil
	case "PLN":
		return Currency{code: "PLN", precision: 2}, nil
	case "RON":
		return Currency{code: "RON", precision: 2}, nil
	case "SEK":
		return Currency{code: "SEK", precision: 2}, nil
	case "CHF":
		return Currency{code: "CHF", precision: 2}, nil
	case "ISK":
		return Currency{code: "ISK", precision: 2}, nil
	case "NOK":
		return Currency{code: "NOK", precision: 2}, nil
	case "HRK":
		return Currency{code: "HRK", precision: 2}, nil
	case "TRY":
		return Currency{code: "TRY", precision: 2}, nil
	case "AUD":
		return Currency{code: "AUD", precision: 2}, nil
	case "BRL":
		return Currency{code: "BRL", precision: 2}, nil
	case "CAD":
		return Currency{code: "CAD", precision: 2}, nil
	case "CNY":
		return Currency{code: "CNY", precision: 2}, nil
	case "HKD":
		return Currency{code: "HKD", precision: 2}, nil
	case "IDR":
		return Currency{code: "IDR", precision: 2}, nil
	case "ILS":
		return Currency{code: "ILS", precision: 2}, nil
	case "INR":
		return Currency{code: "INR", precision: 2}, nil
	case "KRW":
		return Currency{code: "KRW", precision: 2}, nil
	case "MXN":
		return Currency{code: "MXN", precision: 2}, nil
	case "MYR":
		return Currency{code: "MYR", precision: 2}, nil
	case "NZD":
		return Currency{code: "NZD", precision: 2}, nil
	case "PHP":
		return Currency{code: "PHP", precision: 2}, nil
	case "SGD":
		return Currency{code: "SGD", precision: 2}, nil
	case "THB":
		return Currency{code: "THB", precision: 2}, nil
	case "ZAR":
		return Currency{code: "ZAR", precision: 2}, nil
	default:
		return Currency{}, errUnknownCurrency
	}
}
