package palindrome

import (
	"regexp"
	"strconv"
	"strings"
)

func IsPalindrome(s string) bool {
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}

func cleanString(s string) string {
	cleaned := s
	cleaned = strings.ToLower(s)
	cleaned = strings.Join(strings.Fields(cleaned), "") // Remove spaces
	regex, _ := regexp.Compile(`[^\p{L}\p{N} ]+`)       // Regular expression for alphanumeric only characters
	return regex.ReplaceAllString(cleaned, "")
}
