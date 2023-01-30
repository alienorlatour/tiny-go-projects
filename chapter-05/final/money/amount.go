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
