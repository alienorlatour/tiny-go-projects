package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	language := "en"
	expectedGreeting := "Hello world"

	greeting := greet(locale(language))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	language := "fr"
	expectedGreeting := "Bonjour le monde"

	greeting := greet(locale(language))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	language := "akk"
	// Akkadian is not implemented yet!
	expectedGreeting := ""

	greeting := greet(locale(language))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}
