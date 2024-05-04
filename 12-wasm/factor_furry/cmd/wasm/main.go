package main

import (
	"fmt"
	"math/rand"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(generate))
	js.Global().Set("validate", js.FuncOf(validate))
	<-c
}

func validate(this js.Value, args []js.Value) any {
	fmt.Println("args of validate:", args)
	expected := js.ValueOf(args[0]).Int()
	answer := js.ValueOf(args[1]).Int()

	var result bool

	// Comparing with the answer provided by the user
	if expected == answer {
		fmt.Println("Correct!")
		result = true

		generate(this, args)
	} else {
		fmt.Println("Try again!")
	}

	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", result)
	return nil
}

func generate(this js.Value, args []js.Value) any {
	operand1 := rand.Intn(10) + 1
	operand2 := rand.Intn(10) + 1
	expected := operand1 * operand2

	fmt.Printf("%dx%d=%d\n", operand1, operand2, expected)

	js.Global().Set("operand1", operand1)
	js.Global().Set("operand2", operand2)
	js.Global().Set("expected", expected)

	return []interface{}{operand1, operand2, expected}
}
