package money

// exchangeRates is the dependency that fetches applicable rates.
type exchangeRates interface {
}

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	return Amount{}, nil
}
