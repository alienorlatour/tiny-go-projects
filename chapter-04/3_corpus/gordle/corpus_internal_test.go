package gordle

import (
	"testing"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file   string
		length int
		err    error
	}{
		"english corpus": {
			file:   "../corpus/english.txt",
			length: 35,
			err:    nil,
		},
		"empty corpus": {
			file:   "../corpus/empty.txt",
			length: 0,
			err:    errCorpusIsEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			words, err := ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("expected err %v, got %v", tc.err, err)
			}

			if tc.length != len(words) {
				t.Errorf("expected %d, got %d", tc.length, len(words))
			}
		})
	}
}

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
