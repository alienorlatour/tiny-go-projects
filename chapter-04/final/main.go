package main

import (
	_ "embed"
	"errors"
	"os"
	"strings"

	"tiny-go-projects/chapter04/gordle"
)

func readCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	// we expect the corpus to be a line-separated list of words
	words := strings.Split(string(data), "\n")
	if len(words) == 0 {
		return nil, errors.New("corpus is empty")
	}
	return words, nil
}

const maxAttempts = 6

func main() {
	corpus, err := readCorpus("corpus/english.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	// Use the default values for every parameter, but set the default number of max attempts to 6.
	g, err := gordle.New(corpus, gordle.WithMaxAttempts(maxAttempts))
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}
