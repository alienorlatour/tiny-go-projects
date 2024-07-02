package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("validate", js.FuncOf(validate))
	<-c
}

func generate(this js.Value, args []js.Value) any {
	operand1 := rand.IntN(10) + 1
	operand2 := rand.IntN(10) + 1

	js.Global().Get("document").Call("getElementById", "operand1").Set("innerHTML", operand1)
	js.Global().Get("document").Call("getElementById", "operand2").Set("innerHTML", operand2)

	return nil
}

func validate(this js.Value, args []js.Value) any {
	defer func() {
		// Reset the contents of the fields
		js.Global().Get("document").Call("getElementById", "providedAnswer").Set("value", "")
	}()

	// Retrieve the operands.
	operand1 := js.Global().Get("document").Call("getElementById", "operand1").Get("innerHTML")
	op1, err := strconv.Atoi(operand1.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	operand2 := js.Global().Get("document").Call("getElementById", "operand2").Get("innerHTML")
	op2, err := strconv.Atoi(operand2.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	a, err := strconv.Atoi(args[0].String())
	if err != nil {
		js.Global().Call("alert", fmt.Sprintf("not a number: %q", args[0].String()))
		return nil
	}

	fmt.Printf("comparing %dx%d with %d", op1, op2, a)
	// Comparing with the answer provided by the user
	if op1*op2 == a {
		js.Global().Call("alert", "Bravo !")
		generate(this, args)
	} else {
		js.Global().Call("alert", "Try again... "+args[0].String()+" is not the correct answer.")
	}

	return nil
}
