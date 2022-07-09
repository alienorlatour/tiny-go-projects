package gordle

import (
	"strings"
	"testing"
)

func TestCorpus_randomWord(t *testing.T) {
	word := randomWord()

	if !strings.Contains(corpus, string(word)) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}
