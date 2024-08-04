package main

import (
	"math/rand/v2"
	"syscall/js"
)

type multiplication struct {
	opLeft, opRight int
}

func main() {
	m := &multiplication{}

	// Registering Go functions to JavaScript
	js.Global().Set("generate", js.FuncOf(m.generate))

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
