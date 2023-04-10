package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	// preparation phase: define the expected returned value
	expectedGreeting := "Hello world"

	// execution phase: call the examined greet function
	greeting := greet()

	// decision phase: check the returned value
	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}
