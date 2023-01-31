package money

type exchangeRates interface {
	// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}

// ExchangeRate is the rate to convert from a currency to another.
// It is a float64, because the precision of an official change rate is 5 significant digits.
type ExchangeRate float64
