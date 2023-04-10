package main

import "fmt"

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents the language’s code
type language string

// greet says hello to the world in the specified language
func greet(lang language) string {
	switch lang {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
