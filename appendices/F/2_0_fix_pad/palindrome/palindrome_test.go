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
	f.Add("01")   // pads with 0

	// Call Fuzz and verify output
	f.Fuzz(func(t *testing.T, input string) {
		got := IsPalindromeNumber(input)

		// Basic validation: a string reversed should match the palindrome status
		reversed := reverseInt(input)
		want := IsPalindromeNumber(reversed)
		if got != want {
			t.Errorf("Palindrome mismatch for input: %s and its reverse: %s", input, reversed)
		}
	})
}

// reverseInt reverses an integer using the built-in slices.Reverse function
func reverseInt(s string) string {
	split := []rune(s)
	slices.Reverse(split)
	return string(split)
}
