package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ablqk/tiny-go-projects/chapter-04/3_corpus/gordle"
)

func main() {
	fmt.Println("Welcome to Gordle!")

	corpus, err := readCorpus("corpus/english.txt")
	if err != nil {
		panic(err)
	}

	// Create the game.
	g := gordle.New(bufio.NewReader(os.Stdin), corpus, 6)

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
