package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"syscall/js"
)

type multiplication struct {
	opLeft, opRight int
}

func main() {
	m := &multiplication{}

	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(m.generate))
	js.Global().Set("validate", js.FuncOf(m.validate))

	// Wait forever
	<-make(chan struct{})
}

func (m *multiplication) generate(_ js.Value, _ []js.Value) any {
	m.opLeft = rand.IntN(11)
	m.opRight = rand.IntN(11)
	document := js.Global().Get("document")

	document.Call("getElementById", "operand1").Set("innerHTML", m.opLeft)
	document.Call("getElementById", "operand2").Set("innerHTML", m.opRight)

	return nil
}

func (m *multiplication) validate(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return nil
	}

	document := js.Global().Get("document")

	defer func() {
		// Reset the contents of the input field after the user clicked Validate
		document.Call("getElementById", "providedAnswer").Set("value", "")
	}()

	guess := args[0].String()
	numGuess, err := strconv.Atoi(guess)
	if err != nil {
		js.Global().Call("alert", fmt.Sprintf("not a number: %q", guess))
		return nil
	}

	// Comparing with the answer provided by the user
	if m.opLeft*m.opRight == numGuess {
		js.Global().Call("alert", "Bravo! Here's a new exercise.")
		m.generate(this, args)
	} else {
		js.Global().Call("alert", "Try again... "+guess+" is not the correct answer.")
	}

	return nil
}
