package main

import (
	"fmt"

	"learngo-pockets/logger/pocketlog"
)

func main() {
	lvl := pocketlog.LevelDebug
	fmt.Printf("Level: %v\n", lvl)
}
