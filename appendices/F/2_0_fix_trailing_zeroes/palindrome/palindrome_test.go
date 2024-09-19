package palindrome

import (
	"slices"
	"testing"
)

func FuzzIsPalindromeNumber(f *testing.F) {
	// Seed corpus with some basic test cases
	f.Add("1221") // nominal case
	f.Add("")     // empty string
	f.Add("10")   // ends with 0
	f.Add("01")   // starts with 0
	f.Add("-020") // pads with 0
	f.Add("0")    // only 0

	// Fun other test cases
	f.Add("\xff")                // not a number
	f.Add("+1")                  // signed +
	f.Add("-1")                  // signed -
	f.Add("9700000000000000079") // big
	f.Add("1.1")                 // float
	f.Add("1 221")               // with a space
	f.Add(`{"test":1}`)          // JSON
	f.Add("1_221")               // a valid int with "_"

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
