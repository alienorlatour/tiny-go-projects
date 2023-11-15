package habit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_completeHabit(t *testing.T) {
	t.Parallel()

	t.Run("Full", testCompleteHabitFull)
	t.Run("Partial", testCompleteHabitPartial)
	t.Run("SpaceName", testCompleteHabitSpaceName)
}

func testCompleteHabitFull(t *testing.T) {
	t.Parallel()

	h := Habit{
		ID:              "987",
		Name:            "laugh",
		WeeklyFrequency: 256,
		CreationTime:    time.Date(2023, 10, 27, 1, 5, 0, 0, time.UTC),
	}

	got, err := completeHabit(h)
	require.NoError(t, err)
	assert.Equal(t, h, got)
}

func testCompleteHabitPartial(t *testing.T) {
	t.Parallel()

	h := Habit{
		Name:            "laugh",
		WeeklyFrequency: 256,
	}

	got, err := completeHabit(h)
	require.NoError(t, err)
	assert.Equal(t, h.Name, got.Name)
	assert.Equal(t, h.WeeklyFrequency, got.WeeklyFrequency)
	assert.NotEmpty(t, got.ID)
	assert.NotEmpty(t, got.CreationTime)
}

func testCompleteHabitSpaceName(t *testing.T) {
	t.Parallel()

	h := Habit{
		Name:            "    ",
		WeeklyFrequency: 256,
	}

	_, err := completeHabit(h)
	assert.ErrorAs(t, err, &InvalidInputError{})
}
