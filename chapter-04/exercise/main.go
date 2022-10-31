package main

import (
	"github.com/ablqk/tiny-go-projects/chapter-04/exercise/gordle"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("corpus/hindi.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	g, err := gordle.New(corpus, gordle.WithMaxAttempts(maxAttempts))
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}