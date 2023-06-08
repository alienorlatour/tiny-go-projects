package gordle

import (
	"testing"
)

func TestPickRandomWord(t *testing.T) {
	words := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word, err := PickRandomWord(words)
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	if !inCorpus(words, word) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}

func inCorpus(words []string, word string) bool {
	for _, corpusWord := range words {
		if corpusWord == word {
			return true
		}
	}
	return false
}

// OverrideCorpus allows a test to override the corpus in Gordle.
func OverrideCorpus(newCorpus string) {
	corpus = newCorpus
}
