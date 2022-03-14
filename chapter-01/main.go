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
	return dictionary[l]
}
