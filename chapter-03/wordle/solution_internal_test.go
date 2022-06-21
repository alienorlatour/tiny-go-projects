package wordle

import "testing"

func TestSolution_Feedback(t *testing.T) {
	tt := map[string]struct {
		attempt          []byte
		solution         Solution
		expectedFeedback []status
	}{
		"nominal": {
			attempt:          []byte("hertz"),
			solution:         NewSolution([]byte("hertz")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter": {
			attempt:          []byte("hello"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter with wrong answer": {
			attempt:          []byte("helll"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			attempt:          []byte("lllll"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			attempt:          []byte("hlleo"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			attempt:          []byte("hlllo"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			attempt:          []byte("lllww"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped letters": {
			attempt:          []byte("holle"),
			solution:         NewSolution([]byte("hello")),
			expectedFeedback: []status{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent letter": {
			attempt:          []byte("hulfo"),
			solution:         NewSolution([]byte("helfo")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent letter and incorrect": {
			attempt:          []byte("hulpp"),
			solution:         NewSolution([]byte("helpo")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.solution.Feedback(tc.attempt)
			if !compare(tc.expectedFeedback, got) {
				t.Errorf("attempt: %s, got the wrong feedback, expected %v, got %v", tc.attempt, tc.expectedFeedback, got)
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
