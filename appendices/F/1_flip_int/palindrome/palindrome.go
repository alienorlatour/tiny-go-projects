package palindrome

import (
	"strconv"
)

// IsPalindromeNumber returns if an integer is readable in both ways.
// "1221" is a palindrome
func IsPalindromeNumber(s string) bool {
	// convert to a number
	myInt, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	// flip the integer and compare
	myIntOG := myInt
	myInt2 := 0
	for myInt > 0 {
		myInt2 = 10*myInt2 + myInt%10 // 0*10 + 1234%10 = 4
		myInt = myInt / 10            // 1234/10 = 123
	}

	return myIntOG == myInt2
}
