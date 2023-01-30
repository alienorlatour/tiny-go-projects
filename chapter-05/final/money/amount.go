package money

import (
	"fmt"
	"math"
)

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	number   Number
	currency Currency
}

// NewAmount returns a new Amount.
func NewAmount(number Number, currency Currency) Amount {
	return Amount{number: number, currency: currency}
}

// String implements stringer and returns the Number formatted as
// digits optionally a decimal point followed by digits.
func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.number.String(), a.currency.code)
}

const (
	// ErrTooSmall is returned if the converted amount is too small, in order to protect against precision issues.
	ErrTooSmall = moneyError("amount is less than 1.00")
	// ErrTooLarge is returned if the amount is too large - this would cause floating point precision errors.
	ErrTooLarge = moneyError("amount over 10^15 is too large")
	// ErrTooPrecise is returned if the amount is too precise.
	ErrTooPrecise = moneyError("amount value is too precise")

	// maxAmount value is a thousand billion, using the short scale -- 10^12.
	maxAmount = 1e12
)

func (a Amount) validate() error {
	switch {
	case a.number.tooSmall(a.currency):
		return ErrTooSmall
	case a.number.integerPart > maxAmount:
		return ErrTooLarge
	case a.number.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}

// applyChangeRate returns a new Number representing n multiplied by the rate.
// The precision is the same in and out.
func (a Amount) applyChangeRate(rate ExchangeRate, target Currency) Amount {
	converted := a.number.float() * float64(rate)

	floor := math.Floor(converted)
	decimal := math.Round((converted - floor) * math.Pow10(target.precision))

	return Amount{
		number: Number{
			integerPart: int(floor),
			decimalPart: int(decimal),
			precision:   target.precision,
		},
		currency: target,
	}
}
