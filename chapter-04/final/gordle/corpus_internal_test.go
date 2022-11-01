package gordle

import (
	"testing"
)

func TestPickWord(t *testing.T) {
	corpus, err := ReadCorpus("../corpus/english.txt")
	if err != nil {
		t.Errorf("failed to read corpus")
	}
	word := pickWord(corpus)

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
