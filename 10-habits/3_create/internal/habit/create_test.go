package habit_test

import (
	"context"
	"testing"
	"time"

	"learngo-pockets/habits/internal/habit"

	"github.com/stretchr/testify/assert"
)

func TestCreateHabit(t *testing.T) {
	h := habit.Habit{
		Name:            "swim",
		WeeklyFrequency: 2,
		CreationTime:    time.Now(),
		ID:              "123",
	}

	tests := map[string]struct {
		expectedErr error
	}{
		"nominal": {
			expectedErr: nil,
		},
	}

	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := habit.CreateHabit(context.Background(), h)
			assert.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, h, got)
		})
	}
}
