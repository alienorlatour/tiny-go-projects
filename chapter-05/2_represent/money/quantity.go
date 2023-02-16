package money

import (
	"fmt"
	"strconv"
	"strings"
)

// Quantity can represent money with a fixed precision.
// example: 1.52 = 1 + 52 * 10^(-2) will be stored as {1, 52, 2}
type Quantity struct {
	// cents is the amount of money, not necessarily in
	cents int
	// precision of the cents, as the exponent of a power of 10
	exp int
}

const (
	// ErrInvalidValue is returned if the number is malformed.
	ErrInvalidValue = Error("unable to convert the value")

	// ErrTooLarge is returned if the amount is too large - this would cause floating point precision errors.
	ErrTooLarge = Error("amount over 10^15 is too large")

	// max value is a thousand billion, this is the power of 10.
	maxAmountExp = 12
)

// ParseNumber converts a string into its Quantity representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseNumber(value string) (Quantity, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	if len(intPart) > maxAmountExp {
		return Quantity{}, ErrTooLarge
	}

	cents, err := strconv.Atoi(intPart + fracPart)
	if err != nil {
		return Quantity{}, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	precision := len(fracPart)

	return Quantity{cents: cents, exp: precision}, nil
}
