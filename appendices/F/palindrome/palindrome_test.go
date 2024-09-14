package palindrome

import (
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		//{"palindrome with letters", args{"kayak"}, true},
		//{"palindrome with numbers", args{"1221"}, true},
		//{"a weird one", args{" "}, true},
		//{"not a palindrome", args{"hello"}, false},
		//{"not a palindrome", args{"ab"}, false},
		//{"not a palindrome", args{" 0"}, true},
		////{"palindrome with capital", args{"A0a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzIsPalindrome(f *testing.F) {
	// Seed corpus with some basic test cases
	//f.Add("racecar")
	//f.Add("A man, a plan, a canal, Panama")
	//f.Add("hello")
	f.Add("")
	f.Add("1221")
	//f.Add("1 221")

	f.Fuzz(func(t *testing.T, input string) {
		result := IsPalindrome(input)

		// Basic validation: a string reversed should match the palindrome status
		reversed := reverseInt(input)
		if result != IsPalindrome(reversed) {
			t.Errorf("Palindrome mismatch for input: %s and its reverse: %s", input, reversed)
		}
	})
}

func reverseInt(s string) string {
	_, err := strconv.Atoi(s)
	if err != nil {
		return ""
	}

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
