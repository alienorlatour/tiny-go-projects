package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"learngo-pockets/templates/internal/habit"
	"learngo-pockets/templates/internal/handlers/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_index(t *testing.T) {
	rr := httptest.NewRecorder()

	cli := mocks.NewHabitsClientMock(t)
	cli.ListHabitsMock.Set(func(_ context.Context, _ time.Time) ([]habit.Habit, error) {
		return []habit.Habit{
			{Name: "Knit", WeeklyFrequency: 3, Ticks: 2},
			{Name: "Water the plants", WeeklyFrequency: 1, Ticks: 0},
		}, nil
	})

	s := New(cli, t)
	s.index(rr, httptest.NewRequest(http.MethodGet, "/", nil))

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	expect, err := os.ReadFile("testdata/index.html")
	require.NoError(t, err)
	assert.Equal(t, string(expect), rr.Body.String())
}
