package money

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Number is a structure that can hold a Number with a fixed precision.
// example: 1.52 = 1 + 52 * 10^(-2) will be stored as {1, 52, 2}
type Number struct {
	// the integer part of the Number
	integerPart int
	// precision of the decimal part, as the exponent of a power of 10
	precision int
}

const (
	// errInvalidInteger is raised when parsing string into integer.
	errInvalidInteger = moneyError("unable to parse integer")
	errInvalidDecimal = moneyError("unable to convert decimal part")
)

// matchFloat regexp matches with a float
var matchFloat = regexp.MustCompile("^((0|[1-9]\\d*)\\.(\\d*)?|\\.\\d+)$")

// parseNumber converts a string into its Number representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func parseNumber(value string) (Number, error) {
	intPart, decPart, found := strings.Cut(value, ".")

	// validate the integer part
	_, err := strconv.Atoi(intPart)
	if err != nil {
		return Number{}, fmt.Errorf("%w: %s", errInvalidInteger, err.Error())
	}

	// store the integer
	fullValue := intPart

	if found {
		// validate the decimal part
		_, err := strconv.Atoi(decPart)
		if err != nil {
			return Number{}, fmt.Errorf("%w: %s", errInvalidDecimal, err.Error())
		}

		// store the decimal
		fullValue += decPart
	}

	// store the precision
	precision := len(decPart)

	// convert the full value into an integer
	number, err := strconv.Atoi(fullValue)
	if err != nil {
		return Number{}, fmt.Errorf("%w: %s", errInvalidInteger, err.Error())
	}

	return Number{integerPart: number, precision: precision}, nil
}

// applyChangeRate returns a new Number representing n multiplied by the rate.
// The precision is the same in and out.
func (n Number) applyChangeRate(rate ChangeRate, toPrecision int) Number {
	// compute the new amount
	converted := n.float() * float64(rate)
	// shift the period to handle an integer
	inFloat := converted * math.Pow10(toPrecision)

	// TODO: I am not sure if we should round or trunc here?
	// TODO: computed value: 4394.8871619701395, should we go with 43.84 or 43.95
	// TODO: Add a proper comment depending on what we decide
	integerPart := int(math.Round(inFloat))

	return Number{
		integerPart: integerPart,
		precision:   toPrecision,
	}
}

func (n Number) float() float64 {
	return float64(n.integerPart) * math.Pow10(-n.precision)
}

func (n Number) integer() float64 {
	return float64(n.integerPart) * math.Pow10(n.precision)
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
	// ErrInputTooPrecise is returned if the input amount is too precise, which could lead to precision issues.
	ErrInputTooPrecise = moneyError("input amount should be at least 1 * precision")

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
	case n.tooPrecise(sourceCurrency):
		return ErrInputTooPrecise
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
	return n.float() > maxAmount
}

// String implements stringer and returns the Number formatted as
// digits optionally a decimal point followed by digits.
func (n Number) String() string {
	// for a precision of 2 digits formats %.02f
	format := fmt.Sprintf("%%.0%df", n.precision)

	return fmt.Sprintf(format, n.float())
}
