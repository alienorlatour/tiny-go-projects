package palindrome

import (
	"slices"
	"testing"
)

func FuzzIsPalindromeNumber(f *testing.F) {
	// Seed corpus with some basic test cases
	// Seed corpus with some basic test cases
	f.Add("1221")                // nominal case
	f.Add("")                    // empty string
	f.Add("10")                  // ends with 0
	f.Add("01")                  // pads with 0
	f.Add("\xff")                // not a number
	f.Add("+1")                  // signed +
	f.Add("-1")                  // signed -
	f.Add("9700000000000000079") // too big

	// Fun other test cases
	f.Add("1.1")        // float
	f.Add("1 221")      // with a space
	f.Add(`{"test":1}`) // JSON
	f.Add("1_221")      // a valid int with "_"

	// Call Fuzz and verify output
	f.Fuzz(func(t *testing.T, input string) {
		got, err := IsPalindromeNumber(input)
		if err != nil {
			// V2 ignore if it's not a number
			return
		}

		// Basic validation: a string reversed should match the palindrome status
		reversed := reverseInt(input)
		want, err := IsPalindromeNumber(reversed)
		if err != nil {
			// V3 reversed != than input and not a number, e.g.: string("+0")
			t.Errorf("reversed is not a number: %v", err)
		}
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
