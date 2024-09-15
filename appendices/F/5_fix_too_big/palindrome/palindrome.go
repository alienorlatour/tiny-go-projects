package palindrome

import (
	"fmt"
	"strconv"
	"strings"
)

// IsPalindromeNumber returns if an integer is readable in both ways.
// Only unsigned int are supported.
// "1221" is a palindrome
// "01" is not a palindrome
// "-1" is not a palindrome
func IsPalindromeNumber(s string) (bool, error) {
	// value could be bigger than max int e.g.: 10000000000000000000
	if len(s) > 10 {
		return false, fmt.Errorf("input is too long")
	}

	// check if it's a number, e.g.: "ң\xa3" is not a number
	_, err := strconv.Atoi(s)
	if err != nil {
		// returns an error if it's not a number
		return false, fmt.Errorf("not a number: %s", s)
	}

	// check if there is a sign + or -
	if strings.TrimLeft(s, "+-") != s {
		return false, fmt.Errorf("has sign: %s", s)
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, nil
		}
	}
	return true, nil
}
