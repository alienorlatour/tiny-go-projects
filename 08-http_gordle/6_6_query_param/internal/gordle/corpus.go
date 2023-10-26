package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	// ErrInaccessibleCorpus is returned when the corpus can't be loaded.
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
	// ErrEmptyCorpus is returned when the provided corpus is empty.
	ErrEmptyCorpus = corpusError("corpus is empty")
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

// pickRandomWord returns a random word from the corpus
func pickRandomWord(corpus []string) string {
	// rand.Seed is only necessary if your version of Go is before 1.20.
	// It's best not to have it, if you're using go 1.21 or more recent.
	//nolint:staticcheck // Only if you use Go < 1.20.
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(len(corpus))
	// TODO DONIA use crytpoo

	return corpus[index]
}
