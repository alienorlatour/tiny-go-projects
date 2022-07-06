package gordle

import (
	"errors"
	"fmt"
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
	ts := &testScanner{
		contents: string(expected),
	}
	g := &Gordle{scanner: ts, solution: []rune("MUMMY")}

	got := g.ask()

	if string(got) != string(expected) {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func compare[T ~int](lhs, rhs []T) bool {
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

type testScanner struct {
	contents string
	errors   []string
}

func (*testScanner) Scan() bool {
	return true
}

func (ts *testScanner) Text() string {
	return ts.contents
}

func (ts *testScanner) Err() error {
	if len(ts.errors) > 0 {
		err := fmt.Errorf(ts.errors[0])
		ts.errors = ts.errors[1:]
		return err
	}
	return nil
}
