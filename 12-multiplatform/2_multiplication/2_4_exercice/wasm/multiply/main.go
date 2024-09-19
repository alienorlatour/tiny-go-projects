package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"syscall/js"
)

type scoreTracker struct {
	total   int
	correct int
}

func main() {
	st := &scoreTracker{}
	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("validate", js.FuncOf(st.validate))

	// Wait forever
	<-make(chan struct{})
}

func generate(_ js.Value, _ []js.Value) any {
	operand1 := rand.IntN(10) + 1
	operand2 := rand.IntN(10) + 1
	document := js.Global().Get("document")

	document.Call("getElementById", "operand1").Set("innerHTML", operand1)
	document.Call("getElementById", "operand2").Set("innerHTML", operand2)

	return nil
}

func (st *scoreTracker) validate(this js.Value, args []js.Value) any {
	if len(args) == 0 {
		return nil
	}

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

	fmt.Printf("comparing %dx%d with %d", op1, op2, numGuess)
	st.total++
	// Comparing with the answer provided by the user
	if op1*op2 == numGuess {
		st.correct++
		js.Global().Call("alert", fmt.Sprintf("Bravo! Here's a new exercise.\n%d/%d correct so far!", st.correct, st.total))
		generate(this, args)
	} else {
		js.Global().Call("alert",
			fmt.Sprintf("Try again... %s is not the correct answer.\n%d/%d correct so far!", guess, st.correct, st.total))
	}

	return nil
}
