package main

import (
	"flag"
	"fmt"
)

func main() {
	l := flag.String("lang", "en", "The required language, e.g. en, ur...")
	flag.Parse()

	hello := greet(locale(*l))
	fmt.Println(hello)
}

// locale represents a language
type locale string

// TODO: Rename this as `phrasebook`
// dictionary holds greeting for each supported language
var dictionary = map[locale]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}

// greet says hello to the world
func greet(l locale) string {
	msg, ok := dictionary[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return msg
}
