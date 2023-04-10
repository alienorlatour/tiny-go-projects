package main

import (
	"fmt"

	"learngo-pockets/logger/pocketlog"
)

func main() {
	l := pocketlog.Logger{}

	fmt.Printf("Logger: %#v\n", l)
}
