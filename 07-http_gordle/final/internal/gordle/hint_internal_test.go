package gordle

import "testing"

func Test_feedback_String(t *testing.T) {
	testCases := map[string]struct {
		fb   Feedback
		want string
	}{
		"three correct": {
			fb:   Feedback{correctPosition, correctPosition, correctPosition},
			want: "💚💚💚",
		},
		"one of each": {
			fb:   Feedback{correctPosition, wrongPosition, absentCharacter},
			want: "💚🟡⬜️",
		},
		"different order for one of each": {
			fb:   Feedback{wrongPosition, absentCharacter, correctPosition},
			want: "🟡⬜️💚",
		},
		"unknown position": {
			fb:   Feedback{404},
			want: "💔",
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			if got := testCase.fb.String(); got != testCase.want {
				t.Errorf("String() = %v, want %v", got, testCase.want)
			}
		})
	}
}
