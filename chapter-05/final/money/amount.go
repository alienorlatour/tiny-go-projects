package money

import (
	"fmt"
)

const (
	// ErrTooPrecise is returned if the amount is too precise.
	ErrTooPrecise = moneyError("amount value is too precise")
)

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	number   Number
	currency Currency
}

// NewAmount returns a new Amount.
func NewAmount(number Number, currency Currency) (Amount, error) {
	a := Amount{number: number, currency: currency}

	if err := a.validate(); err != nil {
		return Amount{}, err
	}

	return a, nil
}

// String implements stringer and returns the Number formatted as
// digits optionally a decimal point followed by digits.
func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.number.String(), a.currency.code)
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
