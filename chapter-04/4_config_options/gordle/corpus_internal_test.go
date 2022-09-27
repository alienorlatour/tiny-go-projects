package gordle

import (
	"testing"
)

func TestPickWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word := pickWord(corpus)

	if !inCorpus(corpus, word) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}
