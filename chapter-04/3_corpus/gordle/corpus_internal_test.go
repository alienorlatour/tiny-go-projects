package gordle

import (
	"testing"
)

func TestReadCorpus_EnglishCorpus(t *testing.T) {
	tt := map[string]struct {
		file string
		err  error
	}{
		"english corpus": {
			file: "../corpus/english.txt",
			err:  nil,
		},
		"empty corpus": {
			file: "../corpus/empty.txt",
			err:  errCorpusIsEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			_, err := ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("expected err %s, got %v", tc.err, err)
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
