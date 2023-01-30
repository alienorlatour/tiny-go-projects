package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Number is a structure that can hold a Number with a fixed precision.
// example: 1.52 = 1 + 52 * 10^(-2) will be stored as {1, 52, 2}
type Number struct {
	// the integer part of the Number
	integerPart int
	// the decimal part of the Number
	decimalPart int
	// precision of the decimal part, as the exponent of a power of 10
	precision int
}

const (
	// errors that we might face when parsing numbers
	errInvalidInteger = moneyError("unable to convert integer part")
	errInvalidDecimal = moneyError("unable to convert decimal part")
)

// ParseNumber converts a string into its Number representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseNumber(value string) (Number, error) {
	intPart, decPart, found := strings.Cut(value, ".")

	i, err := strconv.Atoi(intPart)
	if err != nil {
		return Number{}, fmt.Errorf("%w: %s", errInvalidInteger, err.Error())
	}

	var d int
	if found {
		d, err = strconv.Atoi(decPart)
		if err != nil {
			return Number{}, fmt.Errorf("%w: %s", errInvalidDecimal, err.Error())
		}
	}

	precision := len(decPart)

	return Number{integerPart: i, decimalPart: d, precision: precision}, nil
}

const (
	// ErrInputTooSmall is returned if the amount to convert is too small, which could lead to precision issues.
	ErrInputTooSmall = moneyError("input amount should be at least 1.00")
	// ErrOutputTooSmall is returned if the converted amount is too small, in order to protect against precision issues.
	ErrOutputTooSmall = moneyError("output amount is less than 1.00")
	// ErrInputTooLarge is returned if the input amount is too large - this would cause floating point precision errors.
	ErrInputTooLarge = moneyError("input amount over 10^15 is too large")
	// ErrOutputTooLarge is returned if the converted amount is too large, to protect against floating point errors.
	ErrOutputTooLarge = moneyError("output amount is too large (over 10^15)")

	// minAmount
	minAmount = 1
	// maxAmount value is a thousand billion, using the short scale -- 10^12.
	maxAmount = 1e12
)

// validateInput returns an error if the given amount is not the bounded range.
func (n Number) validateInput(sourceCurrency Currency) error {
	switch {
	case n.tooSmall(sourceCurrency):
		return ErrInputTooSmall
	case n.tooBig():
		return ErrInputTooLarge
		//case n.tooPrecise(sourceCurrency):
		//	return ErrOutputTooLarge
	}
	return nil
}

func (n Number) tooPrecise(sourceCurrency Currency) bool {
	return n.precision > sourceCurrency.precision
}

// validateOutput returns an error if the converted amount is not the bounded range.
func (n Number) validateOutput(currency Currency) error {
	switch {
	case n.tooSmall(currency):
		return ErrOutputTooSmall
	case n.tooBig():
		return ErrOutputTooLarge
	}
	return nil
}

func (n Number) tooSmall(currency Currency) bool {
	f := n.float()
	return f != 0 && f < math.Pow10(-currency.precision)
}

func (n Number) tooBig() bool {
	return n.integerPart > maxAmount
}

func (n Number) float() float64 {
	f := float64(n.integerPart)
	f += float64(n.decimalPart) * math.Pow10(-n.precision)

	return f
}

// String implements stringer and returns the Number formatted as
// digits optionally a decimal point followed by digits.
func (n Number) String() string {
	// for a precision of 2 digits formats %d.%02d
	format := fmt.Sprintf("%%d.%%0%dd", n.precision)

	return fmt.Sprintf(format, n.integerPart, n.decimalPart)
}
