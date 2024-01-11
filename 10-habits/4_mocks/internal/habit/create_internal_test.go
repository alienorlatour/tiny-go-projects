package habit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_completeHabit(t *testing.T) {
	t.Parallel()

	t.Run("Full", testValidateAndCompleteHabitFull)
	t.Run("Partial", testValidateAndCompleteHabitPartial)
	t.Run("SpaceName", testValidateAndCompleteHabitSpaceName)
}

func testValidateAndCompleteHabitFull(t *testing.T) {
	t.Parallel()

	h := Habit{
		ID:              "987",
		Name:            "laugh",
		WeeklyFrequency: 256,
		CreationTime:    time.Date(2023, 10, 27, 1, 5, 0, 0, time.UTC),
	}

	got, err := validateAndCompleteHabit(h)
	require.NoError(t, err)
	assert.Equal(t, h, got)
}

func testValidateAndCompleteHabitPartial(t *testing.T) {
	t.Parallel()

	h := Habit{
		Name:            "laugh",
		WeeklyFrequency: 256,
	}

	got, err := validateAndCompleteHabit(h)
	require.NoError(t, err)
	assert.Equal(t, h.Name, got.Name)
	assert.Equal(t, h.WeeklyFrequency, got.WeeklyFrequency)
	assert.NotEmpty(t, got.ID)
	assert.NotEmpty(t, got.CreationTime)
}

func testValidateAndCompleteHabitSpaceName(t *testing.T) {
	t.Parallel()

	h := Habit{
		Name:            "    ",
		WeeklyFrequency: 256,
	}

	_, err := validateAndCompleteHabit(h)
	assert.ErrorAs(t, err, &InvalidInputError{})
}
