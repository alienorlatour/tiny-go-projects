package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	lang := "en"
	expectedGreeting := "Hello world"

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	expectedGreeting := "Bonjour le monde"

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	lang := "akk"
	// Akkadian is not implemented yet!
	expectedGreeting := ""

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}
