package gordle

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestGordle_validateAttempt(t *testing.T) {
	g := &Gordle{solution: []rune("SAUNA")}
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

func TestGordle_ask(t *testing.T) {
	expected := []rune("HELLO")
	ts := &testReader{
		reader: strings.NewReader(string(expected)),
	}
	g := &Gordle{reader: ts, solution: []rune("MUMMY")}

	got := g.ask()

	if string(got) != string(expected) {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func compare(lhs, rhs []status) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for index, value := range lhs {
		if value != rhs[index] {
			return false
		}
	}
	return true
}

type testReader struct {
	reader io.Reader
	error  string
}

func (tr *testReader) Read(p []byte) (n int, err error) {
	if len(tr.error) > 0 {
		return 0, fmt.Errorf(tr.error)
	}

	return tr.reader.Read(p)
}
