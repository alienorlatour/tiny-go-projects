package money

import "fmt"

// ParseAmount transforms user-provided amount into a handy Number type.
func ParseAmount(amount string) (Number, error) {
	n, err := parseNumber(amount)
	if err != nil {
		return Number{}, fmt.Errorf("unable to parse amount: %w", err)
	}

	return n, nil
}

// ParseCurrencies transforms user-provided inputs into handy Currency type.
func ParseCurrencies(from, to string) (Currency, Currency, error) {
	sourceCurrency, err := parseCurrency(from)
	if err != nil {
		return Currency{}, Currency{}, fmt.Errorf("unable to parse source currency: %w", err)
	}

	targetCurrency, err := parseCurrency(to)
	if err != nil {
		return Currency{}, Currency{}, fmt.Errorf("unable to parse target currency: %w", err)
	}

	return sourceCurrency, targetCurrency, nil
}
