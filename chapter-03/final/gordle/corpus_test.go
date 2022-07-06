package gordle

import (
	"strings"
	"testing"
)

func TestCorpus_randomWord(t *testing.T) {
	word := randomWord()

	if len(word) != 5 {
		t.Errorf("expected a word of length %d, got %q", 5, word)
	}

	if !strings.Contains(corpus, string(word)) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}
