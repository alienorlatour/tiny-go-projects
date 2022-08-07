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
		reader  *bufio.Reader
		want    []rune
		wantErr bool
	}{
		"5 letters in english": {
			reader:  bufio.NewReader(strings.NewReader("hello")),
			want:    []rune("hello"),
			wantErr: false,
		},
		"5 letters in arabic": {
			reader:  bufio.NewReader(strings.NewReader("مرحبا")),
			want:    []rune("مرحبا"),
			wantErr: false,
		},
		"5 letters in japanese": {
			reader:  bufio.NewReader(strings.NewReader("のうにゅう")),
			want:    []rune("のうにゅう"),
			wantErr: false,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(tc.reader)

			got := g.ask()
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("readRunes() got = %v, want %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGordleValidateAttempt(t *testing.T) {
	g := &Gordle{}
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
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}
