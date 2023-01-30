package money

// Amount represents a given amount of money.
type Amount struct {
	number   Number
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = moneyError("amount value is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(number Number, currency Currency) (Amount, error) {
	if number.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{number: number, currency: currency}, nil
}
