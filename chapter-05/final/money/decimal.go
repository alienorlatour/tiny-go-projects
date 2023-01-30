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

func (n Number) tooSmall(currency Currency) bool {
	f := n.float()
	return f != 0 && f < math.Pow10(-currency.precision)
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
