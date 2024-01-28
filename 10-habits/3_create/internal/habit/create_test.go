package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"learngo-pockets/habits/internal/habit"
)

func TestCreate(t *testing.T) {
	h := habit.Habit{
		Name:            "swim",
		WeeklyFrequency: 2,
		CreationTime:    time.Now(),
		ID:              "123",
	}

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		db          *habitAdder
		expectedErr error
	}{
		"nominal": {
			db:          &habitAdder{},
			expectedErr: nil,
		},
		"error case": {
			db:          &habitAdder{err: dbErr},
			expectedErr: dbErr,
		},
	}

	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := habit.Create(context.Background(), tt.db, h)
			assert.ErrorIs(t, err, tt.expectedErr)
			if tt.expectedErr == nil {
				assert.Equal(t, h.Name, got.Name)
			}
		})
	}
}

// habitAdder is a stub for our database
type habitAdder struct {
	err error
}

// Add implements the habitCreator interface
func (ha *habitAdder) Add(ctx context.Context, habit habit.Habit) error {
	return ha.err
}
