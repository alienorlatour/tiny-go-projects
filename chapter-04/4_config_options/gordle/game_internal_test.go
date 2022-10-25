package gordle

import (
	"bufio"
	"errors"
	"strings"
	"testing"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		reader *bufio.Reader
		want   []rune
	}{
		"5 characters in english": {
			reader: bufio.NewReader(strings.NewReader("HELLO")),
			want:   []rune("HELLO"),
		},
		"5 characters in arabic": {
			reader: bufio.NewReader(strings.NewReader("مرحبا")),
			want:   []rune("مرحبا"),
		},
		"5 characters in japanese": {
			reader: bufio.NewReader(strings.NewReader("こんにちは")),
			want:   []rune("こんにちは"),
		},
		"3 and then 5 characters in japanese": {
			reader: bufio.NewReader(strings.NewReader("こんに\nこんにちは")),
			want:   []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := Game{
				reader:          tc.reader,
				solution:        tc.want,
				solutionChecker: &solutionChecker{solution: tc.want}}

			got := g.ask()
			if !compareRunes(got, tc.want) {
				t.Errorf("readRunes() got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

// compareRunes compares two slices and returns whether they have the same elements.
func compareRunes(s1, s2 []rune) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, v1 := range s1 {
		if v1 != s2[i] {
			return false
		}
	}
	return true
}

func TestGameValidateAttempt(t *testing.T) {
	g := &Game{solution: []rune("hello")}
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("hello"),
			expected: nil,
		},
		"too long": {
			word:     []rune("pocket"),
			expected: errInvalidWordLength,
		},
		"empty": {
			word:     []rune(""),
			expected: errInvalidWordLength,
		},
		"nil": {
			word:     nil,
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			err := g.validateAttempt(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %v, got %v", tc.word, tc.expected, err)
			}
		})
	}
}
