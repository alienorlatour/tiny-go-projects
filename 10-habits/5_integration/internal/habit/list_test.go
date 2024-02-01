package habit_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"learngo-pockets/habits/internal/habit"
	"learngo-pockets/habits/internal/habit/mocks"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestListHabits(t *testing.T) {
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

	dbErr := fmt.Errorf("db unavailable")

	tests := map[string]struct {
		db             func(ctl *minimock.Controller) *mocks.HabitListerMock
		expectedErr    error
		expectedHabits []habit.Habit
	}{
		"empty": {
			db: func(ctl *minimock.Controller) *mocks.HabitListerMock {
				db := mocks.NewHabitListerMock(ctl)
				db.FindAllMock.Expect(ctx).Return(nil, nil)
				return db
			},
			expectedErr:    nil,
			expectedHabits: nil,
		},
		"2 items": {
			db: func(ctl *minimock.Controller) *mocks.HabitListerMock {
				db := mocks.NewHabitListerMock(ctl)
				db.FindAllMock.Expect(ctx).Return(habits, nil)
				return db
			},
			expectedErr:    nil,
			expectedHabits: habits,
		},
		"error case": {
			db: func(ctl *minimock.Controller) *mocks.HabitListerMock {
				db := mocks.NewHabitListerMock(ctl)
				db.FindAllMock.Expect(ctx).Return(nil, dbErr)
				return db
			},
			expectedErr:    dbErr,
			expectedHabits: nil,
		},
	}

	for name, tc := range tests {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := minimock.NewController(t)

			db := tc.db(ctrl)

			got, err := habit.ListHabits(context.Background(), db)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.ElementsMatch(t, tc.expectedHabits, got)
		})
	}
}
