package money

import "fmt"

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
