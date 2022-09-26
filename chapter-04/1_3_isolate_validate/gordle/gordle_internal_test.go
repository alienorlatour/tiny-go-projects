package gordle

import (
	"bufio"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestGordleAsk(t *testing.T) {
	tt := map[string]struct {
		reader *bufio.Reader
		want   []rune
	}{
		"5 characters in english": {
			reader: bufio.NewReader(strings.NewReader("hello")),
			want:   []rune("hello"),
		},
		"5 characters in arabic": {
			reader: bufio.NewReader(strings.NewReader("مرحبا")),
			want:   []rune("مرحبا"),
		},
		"5 characters in japanese": {
			reader: bufio.NewReader(strings.NewReader("こんにちは")),
			want:   []rune("こんにちは"),
		},
		"3 characters in japanese": {
			reader: bufio.NewReader(strings.NewReader("こんに\nこんにちは")),
			want:   []rune("こんにちは"),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := Gordle{reader: tc.reader}

			got := g.ask()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("readRunes() got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGordleValidateAttempt(t *testing.T) {
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
			g := &Gordle{}

			err := g.validateAttempt(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %v, got %v", tc.word, tc.expected, err)
			}
		})
	}
}
