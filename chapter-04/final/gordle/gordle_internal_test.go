package gordle

import (
	"errors"
	"fmt"
	"testing"
)

func TestGordle_checkWord(t *testing.T) {
	tt := map[string]struct {
		attempt          []rune
		solution         []rune
		expectedStatuses []status
	}{
		"nominal": {
			attempt:          []rune("hertz"),
			solution:         []rune("hertz"),
			expectedStatuses: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter": {
			attempt:          []rune("hello"),
			solution:         []rune("hello"),
			expectedStatuses: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter with wrong answer": {
			attempt:          []rune("helll"),
			solution:         []rune("hello"),
			expectedStatuses: []status{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			attempt:          []rune("lllll"),
			solution:         []rune("hello"),
			expectedStatuses: []status{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			attempt:          []rune("hlleo"),
			solution:         []rune("hello"),
			expectedStatuses: []status{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			attempt:          []rune("hlllo"),
			solution:         []rune("hello"),
			expectedStatuses: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			attempt:          []rune("lllww"),
			solution:         []rune("hello"),
			expectedStatuses: []status{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped letters": {
			attempt:          []rune("holle"),
			solution:         []rune("hello"),
			expectedStatuses: []status{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent letter": {
			attempt:          []rune("hulfo"),
			solution:         []rune("helfo"),
			expectedStatuses: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent letter and incorrect": {
			attempt:          []rune("hulpp"),
			solution:         []rune("helpo"),
			expectedStatuses: []status{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, _ := New(WithSolution(tc.solution))
			statuses := g.checkAgainstSolution(tc.attempt)
			if !compare(tc.expectedStatuses, statuses) {
				t.Errorf("attempt: %q, got the wrong feedback, expected %v, got %v", string(tc.attempt), tc.expectedStatuses, statuses)
			}
		})
	}
}

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
