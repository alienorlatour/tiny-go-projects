package main

import (
	"fmt"
	"os"

	"tiny-go-projects/chapter03/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	// Create the game.
	g, err := gordle.New(gordle.WithReader(os.Stdin))
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}
