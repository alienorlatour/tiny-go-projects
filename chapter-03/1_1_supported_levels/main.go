package main

import (
	"fmt"

	"github.com/ablqk/tiny-go-projects/chapter-03/1_1_supported_levels/pocketlog"
)

func main() {
	lvl := pocketlog.LevelDebug
	fmt.Printf("Level: %v\n", lvl)
}
