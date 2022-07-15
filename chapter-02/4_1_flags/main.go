package main

import (
	"flag"
	"fmt"
)

func main() {
	var l string
	flag.StringVar(&l, "lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	greeting := greet(locale(l))
	fmt.Println(greeting)
}

// locale represents a language
type locale string

// phrasebook holds greeting for each supported language
var phrasebook = map[locale]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

// greet says hello to the world
func greet(l locale) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
