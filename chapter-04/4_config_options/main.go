package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ablqk/tiny-go-projects/chapter-04/4_config_options/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	corpus, err := readCorpus("corpus/english.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	g, err := gordle.New(corpus, gordle.WithMaxAttempts(6))
	if err != nil {
		panic(err)
	}

	// Run the game ! It will end when it's over.
	g.Play()
}

func readCorpus(path string) ([]string, error) {
	data, err := ioutil.ReadFile(path)
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
