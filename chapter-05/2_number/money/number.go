package money

import (
	"fmt"
	"strconv"
	"strings"
)

// Number is a structure that can represent moneyhold a Number with a fixed precision.
// example: 1.52 = 1 + 52 * 10^(-2) will be stored as {1, 52, 2}
type Number struct {
	integerPart int
	decimalPart int
	precision   int
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
