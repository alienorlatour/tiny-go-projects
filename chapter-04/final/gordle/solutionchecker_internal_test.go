package gordle

import "testing"

func Test_solutionChecker_check(t *testing.T) {
	tt := map[string]struct {
		attempt          []rune
		sc               *solutionChecker
		expectedStatuses []status
	}{
		"nominal": {
			attempt:          []rune("hertz"),
			sc:               &solutionChecker{solution: []rune("hertz")},
			expectedStatuses: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			attempt:          []rune("hello"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character with wrong answer": {
			attempt:          []rune("helll"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			attempt:          []rune("lllll"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			attempt:          []rune("hlleo"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			attempt:          []rune("hlllo"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			attempt:          []rune("lllww"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped characters": {
			attempt:          []rune("holle"),
			sc:               &solutionChecker{solution: []rune("hello")},
			expectedStatuses: []status{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent character": {
			attempt:          []rune("hulfo"),
			sc:               &solutionChecker{solution: []rune("helfo")},
			expectedStatuses: []status{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent character and incorrect": {
			attempt:          []rune("hulpp"),
			sc:               &solutionChecker{solution: []rune("helpo")},
			expectedStatuses: []status{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {

			statuses := tc.sc.check(tc.attempt)
			if !compareStatus(tc.expectedStatuses, statuses) {
				t.Errorf("attempt: %q, got the wrong feedback, expected %v, got %v", string(tc.attempt), tc.expectedStatuses, statuses)
			}
		})
	}
}

func compareStatus(lhs, rhs []status) bool {
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
