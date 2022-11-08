package gordle

import "testing"

func Test_feedback_String(t *testing.T) {
	testCases := map[string]struct {
		fb   feedback
		want string
	}{
		"three correct": {
			fb:   feedback{correctPosition, correctPosition, correctPosition},
			want: "💚💚💚",
		},
		"one of each": {
			fb:   feedback{correctPosition, wrongPosition, absentCharacter},
			want: "💚🟡⬜️",
		},
		"different order for one of each": {
			fb:   feedback{wrongPosition, absentCharacter, correctPosition},
			want: "🟡⬜️💚",
		},
		"unknown position": {
			fb:   feedback{404},
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
