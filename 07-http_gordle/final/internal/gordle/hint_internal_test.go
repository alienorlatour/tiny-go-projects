package gordle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

func TestFeedbackGameWon(t *testing.T) {
	tt := map[string]struct {
		fb   Feedback
		want bool
	}{
		"game not won": {
			fb:   Feedback{0, 1, 0, 0, 0},
			want: false,
		},
		"game won": {
			fb:   Feedback{2, 2, 2, 2, 2},
			want: true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.fb.GameWon()
			assert.Equal(t, tc.want, got)
		})
	}
}
