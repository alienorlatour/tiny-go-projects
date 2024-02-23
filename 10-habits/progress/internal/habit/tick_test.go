package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/habit/mocks"
)

func TestTickHabit(t *testing.T) {
	ctx := context.Background()

	timestamp := time.Date(2024, time.Month(1), 22, 9, 27, 0, 0, time.UTC)

	h := habit.Habit{
		ID:              "123",
		Name:            "walk",
		WeeklyFrequency: 5,
		CreationTime:    timestamp,
	}

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		habitDB     func(ctl *minimock.Controller) *mocks.HabitFinderMock
		tickDB      func(ctl *minimock.Controller) *mocks.TickAdderMock
		expectedErr error
	}{
		"add tick": {
			habitDB: func(ctl *minimock.Controller) *mocks.HabitFinderMock {
				db := mocks.NewHabitFinderMock(ctl)
				db.FindMock.Expect(ctx, h.ID).Return(h, nil)
				return db
			},
			tickDB: func(ctl *minimock.Controller) *mocks.TickAdderMock {
				db := mocks.NewTickAdderMock(ctl)
				db.AddTickMock.Expect(ctx, h.ID, timestamp).Return(nil)
				return db
			},
			expectedErr: nil,
		},
		"error case on AddTick Tick": {
			habitDB: func(ctl *minimock.Controller) *mocks.HabitFinderMock {
				db := mocks.NewHabitFinderMock(ctl)
				db.FindMock.Expect(ctx, "123").Return(h, nil)
				return db
			},
			tickDB: func(ctl *minimock.Controller) *mocks.TickAdderMock {
				db := mocks.NewTickAdderMock(ctl)
				db.AddTickMock.Expect(ctx, h.ID, timestamp).Return(dbErr)
				return db
			},
			expectedErr: dbErr,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)

			habitDB := tc.habitDB(ctrl)
			tickDB := tc.tickDB(ctrl)

			err := habit.Tick(context.Background(), habitDB, tickDB, h.ID, timestamp)
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
