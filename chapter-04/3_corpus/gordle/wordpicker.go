package gordle

import (
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ReadCorpus reads the file located to the given path
// and returns a list of words
func ReadCorpus(path string) ([]string, error) {
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

// pickWord returns a random word from the corpus
func pickWord(corpus []string) []rune {
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Int() % len(corpus)

	return []rune(strings.ToUpper(corpus[index]))
}
