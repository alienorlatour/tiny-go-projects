package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("validate", js.FuncOf(validate))
	<-c
}

func generate(this js.Value, args []js.Value) any {
	operand1 := rand.Intn(10) + 1
	operand2 := rand.Intn(10) + 1

	fmt.Printf("%dx%d\n", operand1, operand2)

	js.Global().Get("document").Call("getElementById", "operand1").Set("innerHTML", operand1)
	js.Global().Get("document").Call("getElementById", "operand2").Set("innerHTML", operand2)

	return []interface{}{}
}

func validate(this js.Value, args []js.Value) any {
	fmt.Println(args)

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

	// Read the contents as they were provided by the user.
	providedAnswer := js.Global().Get("document").Call("getElementById", "providedAnswer").Get("value")
	if providedAnswer.String() == "" {
		js.Global().Call("alert", "Please give an answer")
		return nil
	}

	a, err := strconv.Atoi(providedAnswer.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	// Comparing with the answer provided by the user
	if op1*op2 == a {
		js.Global().Call("alert", "Bravo !")
		generate(this, args)
	} else {
		js.Global().Call("alert", "Try again...")
	}

	// Reset the contents of the fields
	js.Global().Get("document").Call("getElementById", "providedAnswer").Set("value", "")

	return nil
}
