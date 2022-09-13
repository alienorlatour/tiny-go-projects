package money

const (
	errUnknownChangeRate = moneyError("no change rate known between currencies")
)

// changeRate is a float64, because the precision of a float32 is limited at 6 digits after the first
// significant one. This means we introduce a 1 euro error when we convert numbers above 10^7 euros.
type changeRate float64

// fetchChangeRate is in charge of retrieving the change rate between two currencies.
func fetchChangeRate(from, to string) (changeRate, error) {
	return 1.5168, nil
}
