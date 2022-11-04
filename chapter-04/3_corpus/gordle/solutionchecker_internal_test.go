package gordle

import "testing"

func Test_solutionChecker_check(t *testing.T) {
	tt := map[string]struct {
		guess            []rune
		solution         string
		expectedStatuses feedback
	}{
		"nominal": {
			guess:            []rune("hertz"),
			solution:         "hertz",
			expectedStatuses: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:            []rune("hello"),
			solution:         "hello",
			expectedStatuses: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character with wrong answer": {
			guess:            []rune("helll"),
			solution:         "hello",
			expectedStatuses: feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			guess:            []rune("lllll"),
			solution:         "hello",
			expectedStatuses: feedback{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:            []rune("hlleo"),
			solution:         "hello",
			expectedStatuses: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			guess:            []rune("hlllo"),
			solution:         "hello",
			expectedStatuses: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			guess:            []rune("lllww"),
			solution:         "hello",
			expectedStatuses: feedback{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped characters": {
			guess:            []rune("holle"),
			solution:         "hello",
			expectedStatuses: feedback{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent character": {
			guess:            []rune("hulfo"),
			solution:         "helfo",
			expectedStatuses: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent character and incorrect": {
			guess:            []rune("hulpp"),
			solution:         "helpo",
			expectedStatuses: feedback{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			solChecker := &solutionChecker{solution: []rune(tc.solution)}
			statuses := solChecker.check(tc.guess)
			if !tc.expectedStatuses.Equal(statuses) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got %v", string(tc.guess), tc.expectedStatuses, statuses)
			}
		})
	}
}

// Equal determines equality of two feedbacks.
func (fb feedback) Equal(other feedback) bool {
	if len(fb) != len(other) {
		return false
	}
	for index, value := range fb {
		if value != other[index] {
			return false
		}
	}
	return true
}
