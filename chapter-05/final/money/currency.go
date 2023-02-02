package money

// Currency defines the code of a money and its decimal precision.
type Currency struct {
	code      string
	precision int
	toEuro    float32
}

func (c Currency) Code() string {
	return c.code
}

// errInvalidCurrencyCode is returned when the currency to parse is not a standard 3-letter code.
const errInvalidCurrencyCode = moneyError("invalid currency code")

// ParseCurrency returns the currency associated to a name and may return errUnknownCurrency.
func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, errInvalidCurrencyCode
	}

	switch code {
	case "IRR":
		return Currency{code: code, precision: 0}, nil
	case "MGA", "MRU":
		return Currency{code: code, precision: 1}, nil // the fraction is actually 5
	case "CNY", "VND":
		return Currency{code: code, precision: 1}, nil
	case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
		return Currency{code: code, precision: 3}, nil
	default:
		// All other circulating currencies use a hundredth division.
		return Currency{code: code, precision: 2}, nil
	}
}
