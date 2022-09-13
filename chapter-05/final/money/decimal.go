package money

import (
	"fmt"
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
	toUnit int
}

const (
	// errors that we might face when parsing numbers
	errInvalidInteger = internalError("unable to convert integer part")
	errInvalidDecimal = internalError("unable to convert decimal part")
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

	return number{integerPart: i, decimalPart: d, toUnit: precision}, nil
}
