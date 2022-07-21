package main

import (
	"fmt"

	"tiny-go-projects/chapter03/1_1_supported_levels/pocketlog"
)

func main() {
	lvl := pocketlog.LevelDebug
	fmt.Printf("Level: %v\n", lvl)
}
