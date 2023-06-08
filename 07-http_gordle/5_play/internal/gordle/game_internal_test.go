package gordle

import (
	"errors"
	"testing"

	"golang.org/x/exp/slices"
)

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     string
		expected error
	}{
		"nominal": {
			word:     "GUESS",
			expected: nil,
		},
		"too long": {
			word:     "POCKET",
			expected: ErrInvalidGuessLength,
		},
		"empty": {
			word:     "",
			expected: ErrInvalidGuessLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, _ := New("SLICE")

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%s, expected %v, got %v", tc.word, tc.expected, err)
			}
		})
	}
}

func TestSplitToUppercaseCharacters(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"lowercase": {
			input: "pocket",
			want:  []rune("POCKET"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := splitToUppercaseCharacters(tc.input)

			if !slices.Equal(tc.want, got) {
				t.Errorf("expected %v, got %v", tc.want, got)
			}
		})
	}
}

func Test_computeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback Feedback
	}{
		"nominal": {
			guess:            "HERTZ",
			solution:         "HERTZ",
			expectedFeedback: Feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:            "HELLO",
			solution:         "HELLO",
			expectedFeedback: Feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character with wrong answer": {
			guess:            "HELLL",
			solution:         "HELLO",
			expectedFeedback: Feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentCharacter},
		},
		"five identical, but only two are there": {
			guess:            "LLLLL",
			solution:         "HELLO",
			expectedFeedback: Feedback{absentCharacter, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:            "HLLEO",
			solution:         "HELLO",
			expectedFeedback: Feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
		"three identical, but not in the right position (from left to right)": {
			guess:            "HLLLO",
			solution:         "HELLO",
			expectedFeedback: Feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"one correct, one incorrect, one absent (left of the correct)": {
			guess:            "LLLWW",
			solution:         "HELLO",
			expectedFeedback: Feedback{wrongPosition, absentCharacter, correctPosition, absentCharacter, absentCharacter},
		},
		"swapped characters": {
			guess:            "HOLLE",
			solution:         "HELLO",
			expectedFeedback: Feedback{correctPosition, wrongPosition, correctPosition, correctPosition, wrongPosition},
		},
		"absent character": {
			guess:            "HULFO",
			solution:         "HELFO",
			expectedFeedback: Feedback{correctPosition, absentCharacter, correctPosition, correctPosition, correctPosition},
		},
		"absent character and incorrect": {
			guess:            "HULPP",
			solution:         "HELPO",
			expectedFeedback: Feedback{correctPosition, absentCharacter, correctPosition, correctPosition, absentCharacter},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got the wrong Feedback, expected %v, got %v", tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
