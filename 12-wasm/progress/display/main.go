package main

import (
	"fmt"
	"syscall/js"
)

var (
	generationDone = make(chan struct{})
	solveDone      = make(chan struct{})
)

func main() {
	fmt.Println("Let's solve a maze!")

	/* TODO
	// wraps the Go function generateMaze as a callback and its purpose is the sum operation.
	callbackGenerate := js.FuncOf(generateMaze)
	// free up resources as soon as the function returns
	defer callbackGenerate.Release()

	// setGeneration is the JavaScript property used as a resolver of the promise to generate a maze,
	// that wait for the result of the wrapped Go callback.
	setGeneration := js.Global().Get("setGeneration")
	setGeneration.Invoke(callbackGenerate)
	<-generationDone
	*/

	// wraps the Go function solveMaze as a callback and its purpose to solve the maze.
	callbackSolve := js.FuncOf(solveMaze)
	// free up resources as soon as the function returns
	defer callbackSolve.Release()

	// setResult is the JavaScript property used as a resolver of the promise,
	// that wait for the result of the wrapped Go callback.
	setSolve := js.Global().Get("setSolve")
	setSolve.Invoke(callbackSolve)
	<-solveDone
}

//func generateMaze(Value js.Value, args []js.Value) interface{} {
//
//	generationDone <- struct{}{}
//	return nil
//}

func solveMaze(Value js.Value, args []js.Value) interface{} {
	fmt.Println("Coucou")
	build()

	solveDone <- struct{}{}
	return nil
}
