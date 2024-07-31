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
	m.opLeft = rand.IntN(11)
	m.opRight = rand.IntN(11)

	dom := js.Global().Get("document")

	dom.Call("getElementById", "operand1").Set("innerHTML", m.opLeft)
	dom.Call("getElementById", "operand2").Set("innerHTML", m.opRight)
}
