package money

import "fmt"

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	quantity Quantity
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("amount value is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(quantity Quantity, currency Currency) (Amount, error) {
	if quantity.exp > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{quantity: quantity, currency: currency}, nil
}

func (a Amount) validate() error {
	switch {
	case a.quantity.cents > maxAmount:
		return ErrTooLarge
	case a.quantity.exp > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}

// String implements stringer.
func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.quantity.String(), a.currency.code)
}
