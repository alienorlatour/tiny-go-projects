package money

// exchangeRates is the dependency that fetches applicable rates.
type exchangeRates interface {
}

func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	return Amount{}, nil
}
