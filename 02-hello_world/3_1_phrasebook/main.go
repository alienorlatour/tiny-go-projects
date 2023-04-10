package main

import (
	"fmt"
)

func main() {
	greeting := greet("en")
	fmt.Println(greeting)
}

// language represents a language
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

// greet says hello to the world in various languages
func greet(lang language) string {
	greeting, ok := phrasebook[lang]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", lang)
	}

	return greeting
}
