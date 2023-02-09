package money

import (
	"fmt"
)

// Amount defines a quantity of money in a given currency.
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

func (a Amount) validate() error {
	switch {
	case a.number.integerPart > maxAmount:
		return ErrTooLarge
	case a.number.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}

// String implements stringer.
func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.number.String(), a.currency.code)
}
