package palindrome

import (
	"fmt"
	"strconv"
	"strings"
)

// IsPalindromeNumber returns if an integer is readable in both ways.
// Only unsigned int are supported.
// "1221" is a palindrome
// "-1" is not a palindrome
func IsPalindromeNumber(s string) (bool, error) {
	// value could be bigger than max int e.g.: 9700000000000000000
	if len(s) > 10 {
		return false, fmt.Errorf("input is too long")
	}

	// V1 check if it's a number, e.g.: "ң\xa3"
	_, err := strconv.Atoi(s)
	//myInt, err := strconv.Atoi(s)
	if err != nil {
		// V2 returns an error
		return false, fmt.Errorf("not a number: %s", s)
	}

	// V4
	if strings.TrimLeft(s, "+-") != s {
		return false, fmt.Errorf("has sign: %s", s)
	}

	// V5 left pad with 0 not a palindrome
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false, nil
		}
	}

	// V0 flip the integer and compare
	//myIntOG := myInt
	//myInt2 := 0
	//for myInt > 0 {
	//	myInt2 = 10*myInt2 + myInt%10 // 0*10 + 1234%10 = 4
	//	myInt = myInt / 10            // 1234/10 = 123
	//}

	//return myIntOG == myInt2, nil

	return true, nil
}
