package money

import "fmt"

func Convert(amount, from, to string) (string, error) {
	// parse
	n, err := parseNumber(amount)
	if err != nil {
		return "", fmt.Errorf("unable to parse amount: %w", err)
	}

	// convert
	// format
	return "", nil
}
