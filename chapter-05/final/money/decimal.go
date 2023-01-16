package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// number is a structure that can hold a number with a fixed precision.
// example: 1.52 = 1 + 52 * 10^(-2) will be stored as {1, 52, 2}
type number struct {
	// the integer part of the number
	integerPart int
	// the decimal part of the number
	decimalPart int
	// precision of the decimal part, as the exponent of a power of 10
	precision int
}

const (
	// errors that we might face when parsing numbers
	errInvalidInteger = moneyError("unable to convert integer part")
	errInvalidDecimal = moneyError("unable to convert decimal part")
)

// parseNumber converts a string into its number representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func parseNumber(value string) (number, error) {
	intPart, decPart, found := strings.Cut(value, ".")

	i, err := strconv.Atoi(intPart)
	if err != nil {
		return number{}, fmt.Errorf("%w: %s", errInvalidInteger, err.Error())
	}

	var d int
	if found {
		d, err = strconv.Atoi(decPart)
		if err != nil {
			return number{}, fmt.Errorf("%w: %s", errInvalidDecimal, err.Error())
		}
	}

	precision := len(decPart)

	return number{integerPart: i, decimalPart: d, precision: precision}, nil
}

// applyChangeRate returns a new number representing n multiplied by the rate.
// The precision is the same in and out.
func (n number) applyChangeRate(rate ExchangeRate, toPrecision int) number {
	converted := n.float() * float64(rate)

	floor := math.Floor(converted)
	decimal := math.Round((converted - floor) * math.Pow10(toPrecision))

	return number{
		integerPart: int(floor),
		decimalPart: int(decimal),
		precision:   toPrecision,
	}
}

func (n number) float() float64 {
	f := float64(n.integerPart)
	f += float64(n.decimalPart) * math.Pow10(-n.precision)

	return f
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

	minAmount = 1
	// maxAmount value is a thousand billion, using the short scale -- 10^12.
	maxAmount = 1_000_000_000_000
)

func (n number) validateInput( /*sourceCurrency Currency*/ ) error {
	switch {
	case n.tooSmall():
		return ErrInputTooSmall
	case n.tooBig():
		return ErrInputTooLarge
		//case n.tooPrecise(sourceCurrency):
		//	return ErrOutputTooLarge
	}
	return nil
}

func (n number) tooPrecise(sourceCurrency Currency) bool {
	return n.precision > sourceCurrency.precision
}

func (n number) validateOutput() error {
	switch {
	case n.tooSmall():
		return ErrOutputTooSmall
	case n.tooBig():
		return ErrOutputTooLarge
	}
	return nil
}

func (n number) tooSmall() bool {
	return n.integerPart < minAmount
}

func (n number) tooBig() bool {
	return n.integerPart > maxAmount
}

// String implements stringer and returns the number formatted as
// digits optionally a decimal point followed by digits.
func (n number) String() string {
	// for a precision of 2 digits formats %d.%02d
	format := fmt.Sprintf("%%d.%%0%dd", n.precision)

	return fmt.Sprintf(format, n.integerPart, n.decimalPart)
}
