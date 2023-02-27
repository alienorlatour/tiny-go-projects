package money

import "fmt"

// Amount defines a decimal of money in a given currency.
type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("decimal is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrTooPrecise
	}

	return Amount{quantity: quantity, currency: currency}, nil
}

// validate returns an error if and only if an Amount is unsafe to use.
func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxAmount:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}

// String implements stringer.
func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.quantity.String(), a.currency.code)
}
