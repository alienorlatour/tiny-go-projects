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
	ErrTooLarge = Error("amount over 10^12 is too large")

	// maxAmount value is a thousand billion, using the short scale -- 10^12.
	maxAmount = 1e12
)

// ParseQuantity converts a string into its Quantity representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseQuantity(value string) (Quantity, error) {
	intPart, fracPart, _ := strings.Cut(value, ".")

	cents, err := strconv.Atoi(intPart + fracPart)
	if err != nil {
		return Quantity{}, fmt.Errorf("%w: %s", ErrInvalidValue, err.Error())
	}

	if cents > maxAmount {
		return Quantity{}, ErrTooLarge
	}

	precision := len(fracPart)

	return Quantity{cents: cents, exp: precision}, nil
}

// String implements stringer and returns the Number formatted as
// digits and optionally a decimal point followed by digits.
func (q Quantity) String() string {
	return "" // todo by a better brain...
}
