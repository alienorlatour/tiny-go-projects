package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/habit/mocks"
	"learngo-pockets/habits/internal/tick"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

// TODO FIX ME
func TestTickHabit(t *testing.T) {
	ctx := context.Background()

	habits := []habit.Habit{
		{
			ID:              "123",
			Name:            "walk",
			WeeklyFrequency: 5,
			CreationTime:    time.Now(),
		},
		{
			ID:              "456",
			Name:            "sleep",
			WeeklyFrequency: 7,
			CreationTime:    time.Now(),
		},
	}

	ticks := []tick.Tick{{time.Now()}, {time.Now().Add(5 * time.Minute)}}

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		habitDB     func(ctl *minimock.Controller) *mocks.HabitFinderMock
		tickDB      func(ctl *minimock.Controller) *mocks.TickAdderMock
		expectedErr error
	}{
		"add tick": {
			habitDB: func(ctl *minimock.Controller) *mocks.HabitFinderMock {
				db := mocks.NewHabitFinderMock(ctl)
				db.FindMock.Expect(ctx, "123").Return(habits[0], nil)
				return db
			},
			tickDB: func(ctl *minimock.Controller) *mocks.TickAdderMock {
				db := mocks.NewTickAdderMock(ctl)
				db.AddMock.Expect(ctx, "123", ticks[0], tick.GetISOWeek()).Return(nil)
				return db
			},
			expectedErr: nil,
		},
		"error case on Add Tick": {
			habitDB: func(ctl *minimock.Controller) *mocks.HabitFinderMock {
				db := mocks.NewHabitFinderMock(ctl)
				db.FindMock.Expect(ctx, "123").Return(habits[0], nil)
				return db
			},
			tickDB: func(ctl *minimock.Controller) *mocks.TickAdderMock {
				db := mocks.NewTickAdderMock(ctl)
				//db.AddMock.Expect(ctx, "123", tick.Tick{Timestamp: time.Now()}, tick.GetISOWeek()).Return(dbErr)
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
			defer ctrl.Finish()

			habitDB := tc.habitDB(ctrl)
			tickDB := tc.tickDB(ctrl)

			err := habit.TickHabit(context.Background(), habitDB, tickDB, "123")
			assert.ErrorIs(t, err, tc.expectedErr)
		})
	}
}
