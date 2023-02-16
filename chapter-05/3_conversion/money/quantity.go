package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Quantity can represent money with a fixed precision.
// example: 1.52 = 1 + 52 * 10^(-2) will be stored as {1, 52, 2}
type Quantity struct {
	// the integer part of the Quantity
	integer int
	// the decimal part of the Quantity
	fractional int
	// precision of the decimal part, as the exponent of a power of 10
	precision int
}

const (
	// ErrInvalidInteger is returned if the integer part is not a number.
	ErrInvalidInteger = Error("unable to convert integer part")

	// ErrInvalidDecimal is returned if the decimal part is not a number.
	ErrInvalidDecimal = Error("unable to convert decimal part")

	// ErrTooLarge is returned if the amount is too large - this would cause floating point precision errors.
	ErrTooLarge = Error("amount over 10^15 is too large")

	// maxAmount value is a thousand billion, using the short scale -- 10^12.
	maxAmount = 1e12
)

// ParseNumber converts a string into its Quantity representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseNumber(value string) (Quantity, error) {
	intPart, decPart, found := strings.Cut(value, ".")

	i, err := strconv.Atoi(intPart)
	if err != nil {
		return Quantity{}, fmt.Errorf("%w: %s", ErrInvalidInteger, err.Error())
	}

	if i > maxAmount {
		return Quantity{}, ErrTooLarge
	}

	var d int
	if found {
		d, err = strconv.Atoi(decPart)
		if err != nil {
			return Quantity{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
		}
	}

	precision := len(decPart)

	return Quantity{integer: i, fractional: d, precision: precision}, nil
}

func (n Quantity) float() float64 {
	f := float64(n.integer)
	f += float64(n.fractional) * math.Pow10(-n.precision)

	return f
}
