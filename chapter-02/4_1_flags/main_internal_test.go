package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	type scenario struct {
		language         locale
		expectedGreeting string
	}

	var tests = map[string]scenario{
		"English": {
			language:         "en",
			expectedGreeting: "Hello world",
		},
		"French": {
			language:         "fr",
			expectedGreeting: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			language:         "akk",
			expectedGreeting: `unsupported language: "akk"`,
		},
		"Greek": {
			language:         "el",
			expectedGreeting: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			language:         "he",
			expectedGreeting: "שלום עולם",
		},
		"Urdu": {
			language:         "ur",
			expectedGreeting: "ہیلو دنیا",
		},
		"Vietnamese": {
			language:         "vi",
			expectedGreeting: "Xin chào Thế Giới",
		},
		"Empty": {
			language:         "",
			expectedGreeting: `unsupported language: ""`,
		},
	}

	// range over all the scenarios
	for scenarioName, tc := range tests {
		t.Run(scenarioName, func(t *testing.T) {
			greeting := greet(tc.language)

			if greeting != tc.expectedGreeting {
				t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
			}
		})
	}
}
