package gordle

import (
	"testing"
)

func TestCorpus_randomWord(t *testing.T) {
	corpus := []string{"HELLO", "ხალხი"}
	word := pickWord([]string{"HELLO", ""})

	if !inCorpus(corpus, string(word)) {
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
