package main

import (
	"fmt"
)

func main() {
	hello := greet("en")
	fmt.Println(hello)
}

type locale string

// dictionary holds greeting for each supported language
var dictionary = map[locale]string{
	"en": "Hello, world!",
	"fr": "Bonjour le monde!",
}

// greet says hello to the world
func greet(l locale) string {
	msg, ok := dictionary[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return msg
}
