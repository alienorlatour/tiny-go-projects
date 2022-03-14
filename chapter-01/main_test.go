package main

import "testing"

func Example_main() {
	main()
	// Output: Hello, world!
}

func Test_greet(t *testing.T) {
	msg := greet("en")
	if msg != "Hello, world!" {
		t.Errorf("expected: Hello, world!, got: %s", msg)
	}

	msg = greet("fr")
	if msg != "Bonjour le monde!" {
		t.Errorf("expected: Bonjour le monde!, got: %s", msg)
	}
}
