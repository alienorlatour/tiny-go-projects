package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

var done = make(chan struct{})

func main() {
	// wraps the Go function printResult as a callback and its purpose is the sum operation.
	callback := js.FuncOf(printResult)
	// free up resources as soon as the function returns
	defer callback.Release()

	// setResult is the JavaScript property used as a resolver of the promise,
	// that wait for the result of the wrapped Go callback.
	setResult := js.Global().Get("setResult")
	setResult.Invoke(callback)
	<-done
}

func printResult(value js.Value, args []js.Value) interface{} {
	value1 := args[0].String()
	v1, err := strconv.Atoi(value1)
	if err != nil {
		fmt.Errorf("error %s", err.Error())
		return err
	}
	value2 := args[1].String()
	v2, err := strconv.Atoi(value2)
	if err != nil {
		fmt.Errorf("error %s", err.Error())
		return err
	}

	fmt.Printf("%d\n", v1+v2)
	done <- struct{}{}
	return nil
}
