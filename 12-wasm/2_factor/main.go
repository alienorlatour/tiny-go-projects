//go:build wasm

package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"syscall/js"
)

func main() {
	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("validate", js.FuncOf(validate))

	// Wait forever.
	<-make(chan struct{})
}

func generate(_ js.Value, _ []js.Value) any {
	operand1 := rand.IntN(11)
	operand2 := rand.IntN(11)
	dom := js.Global().Get("document")

	dom.Call("getElementById", "operand1").Set("innerHTML", operand1)
	dom.Call("getElementById", "operand2").Set("innerHTML", operand2)

	return nil
}

func validate(this js.Value, args []js.Value) any {
	dom := js.Global().Get("document")

	defer func() {
		// Reset the contents of the input field after the user clicked Validate
		dom.Call("getElementById", "providedAnswer").Set("value", "")
	}()

	// Retrieve the operands.
	operand1 := dom.Call("getElementById", "operand1").Get("innerHTML")
	op1, err := strconv.Atoi(operand1.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	operand2 := dom.Call("getElementById", "operand2").Get("innerHTML")
	op2, err := strconv.Atoi(operand2.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	guess := args[0].String()
	numGuess, err := strconv.Atoi(guess)
	if err != nil {
		js.Global().Call("alert", fmt.Sprintf("not a number: %q", guess))
		return nil
	}

	// Comparing with the answer provided by the user
	if op1*op2 == numGuess {
		js.Global().Call("alert", "Bravo! Here's a new exercise.")
		generate(this, args)
	} else {
		js.Global().Call("alert", "Try again... "+guess+" is not the correct answer.")
	}

	return nil
}
