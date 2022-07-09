package main

import (
	"fmt"

	"tiny-go-projects/chapter04/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	// Create the game.
	// Use the default values for every parameter, but set the default number of max attempts to 6.
	g, err := gordle.New(gordle.WithMaxAttempts(6))
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}
