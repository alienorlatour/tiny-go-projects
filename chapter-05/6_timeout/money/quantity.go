package money

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Quantity can represent money with a fixed precision.
// example: 1.52 = 152 * 10^(-2) will be stored as {152, 2}
type Quantity struct {
	// cents is the amount of money, not necessarily in hundredths of the unit
	cents int
	// Number of "cents" in a unit, expressed as a power of 10.
	precisionExp int
}

const (
	// ErrInvalidValue is returned if the quantity is malformed.
	ErrInvalidValue = Error("unable to convert the value")

	// ErrTooLarge is returned if the quantity is too large - this would cause floating point precision errors.
	ErrTooLarge = Error("quantity over 10^12 is too large")

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

	return Quantity{cents: cents, precisionExp: precision}, nil
}

// String implements stringer and returns the Quantity formatted as
// digits and optionally a decimal point followed by digits.
func (q Quantity) String() string {
	// Quick-win, no need to do maths.
	if q.precisionExp == 0 {
		return fmt.Sprintf("%d", q.cents)
	}

	centsPerUnit := tenToThe(q.precisionExp)
	frac := q.cents % centsPerUnit
	integer := q.cents / centsPerUnit

	// We always want to print the correct number of digits - even if they finish with 0.
	quantityFormat := fmt.Sprintf("%%d.%%0%dd", q.precisionExp)
	return fmt.Sprintf(quantityFormat, integer, frac)
}

// tenToThe is a quick implementation of how to raise 10 to a given power.
// It's optimised for small powers, and slow for unusually high powers.
func tenToThe(power int) int {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int(math.Pow(10, float64(power)))
	}
}
