package main

import (
	"bytes"
	"errors"
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
			err := validate(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("expected %q, got %q", tc.expected, err)
			}
		})
	}
}

func Test_input(t *testing.T) {
	expected := []byte("hello")
	reader := testReader{
		line: expected,
	}

	got := input(reader)

	if bytes.Equal(got, expected) {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

type testReader struct {
	line []byte
}

func Test_feedback(t *testing.T) {
	tt := map[string]struct {
		attempt          []byte
		solution         solution
		expectedFeedback []status
	}{
		"nominal": {
			attempt:          []byte("hertz"),
			solution:         newSolution([]byte("hertz")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter": {
			attempt:          []byte("hello"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter with wrong answer": {
			attempt:          []byte("helll"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			attempt:          []byte("lllll"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			attempt:          []byte("hlleo"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			attempt:          []byte("hlllo"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			attempt:          []byte("lllww"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped letters": {
			attempt:          []byte("holle"),
			solution:         newSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent letter": {
			attempt:          []byte("hulfo"),
			solution:         newSolution([]byte("helfo")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent letter and incorrect": {
			attempt:          []byte("hulpp"),
			solution:         newSolution([]byte("helpo")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := feedback(tc.attempt, tc.solution)
			if !compare(tc.expectedFeedback, got) {
				t.Errorf("got the wrong solution, expected %v, got %v", tc.expectedFeedback, got)
			}
		})
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

func (tr testReader) ReadLine() (line []byte, isPrefix bool, err error) {
	return tr.line, false, nil
}
