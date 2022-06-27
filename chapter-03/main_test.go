package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func Test_validate(t *testing.T) {
	tt := map[string]struct {
		word     []byte
		expected error
	}{
		"nominal": {
			word:     []byte(`hello`),
			expected: nil,
		},
		"too long": {
			word:     []byte(`pocket`),
			expected: errInvalidWordLength,
		},
		"empty": {
			word:     []byte(``),
			expected: errInvalidWordLength,
		},
		"nil": {
			word:     nil,
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := validateInput(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%s, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func Test_askWord(t *testing.T) {
	expected := []byte("HELLO")
	reader := testReader{
		line: expected,
	}

	got := askWord(reader)

	if !bytes.Equal(got, expected) {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func Test_pickOne(t *testing.T) {
	word := pickOne(corpus)

	if len(word) != wordLength {
		t.Errorf("expected a word of length %d, got %q", wordLength, word)
	}

	if !strings.Contains(corpus, string(word)) {
		t.Errorf("expected a word in the corpus, got %q", word)
	}
}

type testReader struct {
	line []byte
}

func (tr testReader) ReadLine() (line []byte, isPrefix bool, err error) {
	return tr.line, false, nil
}
