package palindrome

// IsPalindromeNumber returns if an integer is readable in both ways.
// "1221" is a palindrome
func IsPalindromeNumber(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
