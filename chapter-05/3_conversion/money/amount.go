package money

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	quantity Quantity
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("quantity is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(quantity Quantity, currency Currency) (Amount, error) {
	if quantity.precisionExp > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{quantity: quantity, currency: currency}, nil
}

// validate returns an error if and only if an Amount is unsafe to use.
func (a Amount) validate() error {
	switch {
	case a.quantity.cents > maxQuantityExp:
		return ErrTooLarge
	case a.quantity.precisionExp > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}
