package money

import "fmt"

// Convert converts an amount in the from currency into the to currency with the targetPrecision number of digits.
func Convert(amount, from, to string) (string, error) {
	// parse
	n, err := parseNumber(amount)
	if err != nil {
		return "", fmt.Errorf("unable to parse amount: %w", err)
	}

	c := currency{
		code:      to,
		precision: 2,
	}

	// get the change rate
	r, err := fetchChangeRate(from, c.code)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errUnknownChangeRate, err.Error())
	}

	// convert
	convertedValue := n.applyChangeRate(r, c.precision)

	// format
	return convertedValue.String(), nil
}
