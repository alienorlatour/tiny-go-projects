package money

import (
	"fmt"
	"math"
)

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	Number   Number
	Currency Currency
}

// NewAmount returns a new Amount.
func NewAmount(number Number, currency Currency) Amount {
	return Amount{Number: number, Currency: currency}
}

// String implements stringer and returns the Number formatted as
// digits optionally a decimal point followed by digits.
func (a Amount) String() string {
	return fmt.Sprintf("%s %s", a.Number.String(), a.Currency.code)
}

// applyChangeRate returns a new Number representing n multiplied by the rate.
// The precision is the same in and out.
func (a Amount) applyChangeRate(rate ExchangeRate, target Currency) Amount {
	converted := a.Number.float() * float64(rate)

	floor := math.Floor(converted)
	decimal := math.Round((converted - floor) * math.Pow10(target.precision))

	return Amount{
		Number: Number{
			integerPart: int(floor),
			decimalPart: int(decimal),
			precision:   target.precision,
		},
		Currency: target,
	}
}
