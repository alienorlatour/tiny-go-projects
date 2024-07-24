//go:build wasm

package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"syscall/js"
)

type multiplication struct {
	opLeft, opRight int
	successes       int
}

func main() {
	product := &multiplication{}

	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(product.generate))
	js.Global().Set("validate", js.FuncOf(product.validate))

	// Wait forever.
	<-make(chan struct{})
}

func (m *multiplication) generate(_ js.Value, _ []js.Value) any {
	m.opLeft = rand.IntN(11)
	m.opRight = rand.IntN(11)

	dom := js.Global().Get("document")

	dom.Call("getElementById", "operand1").Set("innerHTML", m.opLeft)
	dom.Call("getElementById", "operand2").Set("innerHTML", m.opRight)

	return nil
}

func (m *multiplication) validate(this js.Value, args []js.Value) any {
	dom := js.Global().Get("document")

	defer func() {
		// Reset the contents of the input field after the user clicked Validate
		dom.Call("getElementById", "providedAnswer").Set("value", "")
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
		m.successes++
		m.generate(this, args)
	} else {
		js.Global().Call("alert", "Try again... "+guess+" is not the correct answer.")
	}

	return nil
}
