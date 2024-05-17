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

	js.Global().Set("operand1", operand1)
	js.Global().Set("operand2", operand2)

	return []interface{}{}
}

func validate(this js.Value, args []js.Value) any {
	fmt.Println(args)
	//op1 := js.ValueOf(args[0]).Int()
	//fmt.Println(op1)
	//op2 := js.ValueOf(args[1]).Int()
	//fmt.Println(op2)
	//value := js.ValueOf(args[2]).Int()
	//fmt.Println(value)

	operand1 := js.ValueOf(js.Global().Get("document").Call("getElementById", "operand1").Get("innerHTML"))
	op1, err := strconv.Atoi(operand1.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	fmt.Println(js.ValueOf(js.Global().Get("document").Call("getElementById", "operand1").Get("innerHTML")))

	fmt.Println("operand2")
	operand2 := js.ValueOf(js.Global().Get("document").Call("getElementById", "operand2").Get("innerHTML"))
	op2, err := strconv.Atoi(operand2.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}

	answer := js.ValueOf(js.Global().Get("document").Call("getElementById", "answer").Get("innerHTML"))
	a, err := strconv.Atoi(answer.String())
	if err != nil {
		return fmt.Errorf("unknown format: %w", err)
	}
	//op1 := js.ValueOf(args[0]).Int()
	expected := op1 * op2

	var result bool

	// Comparing with the answer provided by the user
	if expected == a {
		fmt.Println("Correct!")
		result = true

		generate(this, args)
	} else {
		fmt.Println("Try again!")
	}

	js.Global().Get("document").Call("getElementById", "result").Set("innerHTML", result)
	return nil
}
