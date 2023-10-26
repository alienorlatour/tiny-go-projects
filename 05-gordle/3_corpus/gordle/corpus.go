package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ErrCorpusIsEmpty is returned when we cannot find any valid solution in the given corpus.
const ErrCorpusIsEmpty = corpusError("corpus is empty")

// ReadCorpus reads the file located at the given path
// and returns a list of words.
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// we expect the corpus to be a line- or space-separated list of words
	words := strings.Fields(string(data))

	return words, nil
}

// pickWord returns a random word from the corpus
func pickWord(corpus []string) string {
	// rand.Seed is only necessary if your version of Go is before 1.20.
	// It's best not to have it, if you're using go 1.21 or more recent.
	//nolint:staticcheck // Still present to explain how things used to be
	rand.Seed(time.Now().UTC().UnixNano())
	index := rand.Intn(len(corpus))

	return corpus[index]
}
