package main

import "testing"

func Example_main() {
	main()
	// Output: Hello, world!
}

func Test_greet(t *testing.T) {
	var tests = map[string]struct {
		language       locale
		wantedGreeting string
	}{
		"English": {
			language:       "en",
			wantedGreeting: "Hello, world!",
		},
		"Français": {
			language:       "fr",
			wantedGreeting: "Bonjour le monde!",
		},
		"Unsupported": {
			language:       "sw",
			wantedGreeting: `unsupported language: "sw"`,
		},
		"Empty": {
			language:       "",
			wantedGreeting: `unsupported language: ""`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			msg := greet(tc.language)
			if msg != tc.wantedGreeting {
				t.Errorf(`expected: %q, got: %q`, tc.wantedGreeting, msg)
			}
		})
	}
}
