package money

import "fmt"

// Convert converts an amount in the from currency into the to currency with the targetPrecision number of digits.
func Convert(amount, from, to string, targetPrecision int) (string, error) {
	// parse
	n, err := parseNumber(amount)
	if err != nil {
		return "", fmt.Errorf("unable to parse amount: %w", err)
	}

	// get the change rate
	r, err := fetchChangeRate(from, to)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errUnknownChangeRate, err.Error())
	}
	// convert
	convertedValue := n.applyChangeRate(r, targetPrecision)

	// format
	return convertedValue.String(), nil
}
