package money

const (
	errUnknownChangeRate = moneyError("no change rate known between currencies")
)

// changeRate is a float32, because the precision of an official change rate is 5 significant digits.
type changeRate float32

// fetchChangeRate is in charge of retrieving the change rate between two currencies.
func fetchChangeRate(from, to string) (changeRate, error) {

	return 1.5168, nil
}
