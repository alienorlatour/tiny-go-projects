package main

import (
	"fmt"
)

func main() {
	hello := greet("en")
	fmt.Println(hello)
}

type locale string

// greet says hello to the world
func greet(locale locale) string {
	switch locale {
	case "en":
		return "Hello, world!"
	case "fr":
		return "Bonjour le monde!"
	default:
		return ""
	}
}
