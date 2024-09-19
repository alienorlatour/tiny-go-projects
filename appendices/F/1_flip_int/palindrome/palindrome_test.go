package palindrome

import (
	"slices"
	"testing"
)

func FuzzIsPalindromeNumber(f *testing.F) {
	// Seed corpus with some basic test cases
	f.Add("1221") // nominal case
	f.Add("")     // empty string: we add an edge case that checks that our code is "soild" enough.

	f.Fuzz(func(t *testing.T, input string) {
		got := IsPalindromeNumber(input)

		// Basic validation: a string reversed should match the palindrome status
		reversed := reverse(input)
		reverseIsPalindrome := IsPalindromeNumber(reversed)
		if got != reverseIsPalindrome {
			t.Errorf("Palindrome mismatch for input: %s (isPalindrome: %t) "+
				"and its reverse: %s (isPalindrome: %t)", input, got, reversed, reverseIsPalindrome)
		}
	})
}

// reverse reverses a string using the standard slices.Reverse function
func reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}
