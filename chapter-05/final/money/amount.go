package money

import "fmt"

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
