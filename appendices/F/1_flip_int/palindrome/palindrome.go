package palindrome

import (
	"strconv"
)

// IsPalindromeNumber returns whether an integer is identical when read both ways.
// "1221" is a palindrome
func IsPalindromeNumber(s string) bool {
	// convert to a number
	toInt, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	// flip the integer and compare
	original := toInt // 1234
	flip := 0
	for toInt > 0 {
		flip = 10*flip + toInt%10 // 4   | 43 | 432 | 4321
		toInt = toInt / 10        // 123 | 12 |   1 |    0
	}

	return original == flip
}
