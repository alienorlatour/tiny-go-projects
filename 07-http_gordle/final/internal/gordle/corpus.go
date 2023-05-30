package gordle

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"
)

const (
	// ErrInaccessibleCorpus is returned when the corpus can't be loaded.
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
	// ErrEmptyCorpus is returned when the provided corpus is empty.
	ErrEmptyCorpus = corpusError("corpus is empty")
	// ErrPickWord is returned when a word has not been picked from the corpus.
	ErrPickWord = corpusError("failed to pick solution")
)

// ReadCorpus reads the file located at the given path
// and returns a list of words. If it fails, the error is ErrInaccessibleCorpus or ErrEmptyCorpus.
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading (%s): %w", path, err, ErrInaccessibleCorpus)
	}

	// we expect the corpus to be a line- or space-separated list of words
	words := strings.Fields(string(data))

	if len(words) == 0 {
		return nil, ErrEmptyCorpus
	}

	return words, nil
}

// PickRandomWord returns a random word from the corpus.
func PickRandomWord(corpus []string) (string, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(corpus))))
	if err != nil {
		return "", fmt.Errorf("failed to rand index (%s): %w", err, ErrPickWord)
	}

	return corpus[index.Int64()], nil
}
