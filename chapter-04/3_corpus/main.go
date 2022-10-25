package main

import (
	"os"

	"github.com/ablqk/tiny-go-projects/chapter-04/3_corpus/gordle"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	g, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}
