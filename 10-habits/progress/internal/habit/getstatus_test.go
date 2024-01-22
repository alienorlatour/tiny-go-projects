package habit_test

import (
	"context"
	"fmt"
	"learngo-pockets/habits/internal/isoweek"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/habit/mocks"
)

func TestGetStatus(t *testing.T) {
	ctx := context.Background()

	h := habit.Habit{
		ID:              "123",
		Name:            "walk",
		WeeklyFrequency: 5,
		CreationTime:    time.Now(),
	}

	timestamp := time.Date(2024, time.Month(2), 21, 3, 2, 2, 2, time.UTC)
	ticks := []time.Time{timestamp, timestamp.Add(5 * time.Minute)}
	isoWeek := isoweek.ISO8601{Week: 3, Year: 2024}

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		habitDB            func(ctl *minimock.Controller) *mocks.HabitFinderMock
		tickDB             func(ctl *minimock.Controller) *mocks.TickFinderMock
		expectedHabit      habit.Habit
		expectedTicksCount int
		expectedErr        error
	}{
		"2 ticks": {
			habitDB: func(ctl *minimock.Controller) *mocks.HabitFinderMock {
				db := mocks.NewHabitFinderMock(ctl)
				db.FindMock.Expect(ctx, "123").Return(h, nil)
				return db
			},
			tickDB: func(ctl *minimock.Controller) *mocks.TickFinderMock {
				db := mocks.NewTickFinderMock(ctl)
				db.FindWeeklyTicksMock.Expect(ctx, "123", isoWeek).Return(ticks, nil)
				return db
			},
			expectedHabit:      h,
			expectedTicksCount: 2,
			expectedErr:        nil,
		},
		"error case on FindWeeklyTicks": {
			habitDB: func(ctl *minimock.Controller) *mocks.HabitFinderMock {
				db := mocks.NewHabitFinderMock(ctl)
				db.FindMock.Expect(ctx, "123").Return(h, nil)
				return db
			},
			tickDB: func(ctl *minimock.Controller) *mocks.TickFinderMock {
				db := mocks.NewTickFinderMock(ctl)
				db.FindWeeklyTicksMock.Expect(ctx, "123", isoWeek).Return(nil, dbErr)
				return db
			},
			expectedErr:   dbErr,
			expectedHabit: habit.Habit{},
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

			h, c, err := habit.GetStatus(context.Background(), habitDB, tickDB, h.ID, isoWeek)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.expectedHabit, h)
			assert.Equal(t, tc.expectedTicksCount, c)
		})
	}
}
