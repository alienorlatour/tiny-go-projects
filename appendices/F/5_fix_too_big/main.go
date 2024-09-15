package _

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	cleaned := cleanString(s)

	// If the cleaned string is empty or has just one character, it's trivially a palindrome
	if len(cleaned) <= 1 {
		return true
	}

	// Check if the cleaned string is a palindrome
	for i := 0; i < len(cleaned)/2; i++ {
		if cleaned[i] != cleaned[len(cleaned)-1-i] {
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
