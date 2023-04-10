package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	lang := "en"
	want := "Hello world"

	got := greet(language(lang))

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	want := "Bonjour le monde"

	got := greet(language(lang))

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}

func TestGreet_Akkadian(t *testing.T) {
	lang := "akk"
	// Akkadian is not implemented yet!
	want := ""

	got := greet(language(lang))

	if got != want {
		t.Errorf("expected: %q, got: %q", want, got)
	}
}
