package gordle

import (
	"testing"
)

func TestPickRandomWord(t *testing.T) {
	corpus := []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"}
	word, err := PickRandomWord(corpus)

	if err != nil {
		t.Errorf("no error expected, %s", err.Error())
	}

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
