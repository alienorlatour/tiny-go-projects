package money

// Amount defines a quantity of money in a given currency.
type Amount struct {
	number   Quantity
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("amount value is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(number Quantity, currency Currency) (Amount, error) {
	if number.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{number: number, currency: currency}, nil
}

func (a Amount) validate() error {
	switch {
	case a.number.integerPart > maxAmount:
		return ErrTooLarge
	case a.number.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}
