package palindrome

import (
	"strconv"
)

// IsPalindromeNumber returns if an integer is readable in both ways.
// "1221" is a palindrome
func IsPalindromeNumber(s string) bool {
	// check if it's a number, e.g.: "kayak" is not a number
	_, err := strconv.Atoi(s)
	if err != nil {
		// returns an error if it's not a number
		return false
	}

	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
