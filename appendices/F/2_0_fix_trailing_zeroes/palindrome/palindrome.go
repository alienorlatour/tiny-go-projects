package palindrome

import (
	"regexp"
)

// IsPalindromeNumber returns whether an integer is identical when read both ways.
// "1221" is a palindrome
// "01" is not a palindrome
// "004400" is a palindrome
func IsPalindromeNumber(s string) bool {
	// input must have only digits
	re := regexp.MustCompile("[0-9]*")
	if !re.MatchString(s) {
		return false
	}

	// compare first and last digits
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}

	return true
}
