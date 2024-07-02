package main

import (
	"math/rand/v2"
	"syscall/js"
)

func main() {
	c := make(chan struct{})
	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(generate))
	<-c
}

func generate(this js.Value, args []js.Value) any {
	operand1 := rand.IntN(10) + 1
	operand2 := rand.IntN(10) + 1

	js.Global().Get("document").Call("getElementById", "operand1").Set("innerHTML", operand1)
	js.Global().Get("document").Call("getElementById", "operand2").Set("innerHTML", operand2)

	return nil
}
