package money

// NewCurrency is a helper for testing without exposing internal fields from Currency.
func NewCurrency(code string, precision int, toEuro float32) Currency {
	return Currency{
		code:      code,
		precision: precision,
		toEuro:    toEuro,
	}
}
