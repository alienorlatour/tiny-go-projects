package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// locale is used as the language’s abbreviation
type locale string

// greet says hello to the world in the specified language
func greet(l locale) string {
	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
