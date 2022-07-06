package main

import (
	"errors"
	"strings"
	"testing"
)

func Test_validate(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune(`hello`),
			expected: nil,
		},
		"too long": {
			word:     []rune(`pocket`),
			expected: errInvalidWordLength,
		},
		"empty": {
			word:     []rune(``),
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
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func Test_askWord(t *testing.T) {
	expected := []rune("HELLO")
	reader := testReader{
		line: expected,
	}

	got := askWord(reader)

	if string(got) != string(expected) {
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
	line []rune
}

func (tr testReader) ReadLine() ([]rune, error) {
	return tr.line, nil
}
