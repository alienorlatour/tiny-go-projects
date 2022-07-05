package gordle

import "testing"

func TestSolution_Feedback(t *testing.T) {
	tt := map[string]struct {
		attempt          []rune
		solution         Solution
		expectedFeedback []status
	}{
		"nominal": {
			attempt:          []rune("hertz"),
			solution:         NewSolution([]rune("hertz")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter": {
			attempt:          []rune("hello"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double letter with wrong answer": {
			attempt:          []rune("helll"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			attempt:          []rune("lllll"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			attempt:          []rune("hlleo"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			attempt:          []rune("hlllo"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			attempt:          []rune("lllww"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped letters": {
			attempt:          []rune("holle"),
			solution:         NewSolution([]rune("hello")),
			expectedFeedback: []status{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent letter": {
			attempt:          []rune("hulfo"),
			solution:         NewSolution([]rune("helfo")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent letter and incorrect": {
			attempt:          []rune("hulpp"),
			solution:         NewSolution([]rune("helpo")),
			expectedFeedback: []status{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.solution.Feedback(tc.attempt)
			if !compare(tc.expectedFeedback, got) {
				t.Errorf("attempt: %c, got the wrong feedback, expected %v, got %v", tc.attempt, tc.expectedFeedback, got)
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
