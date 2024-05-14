package handlers

import (
	"context"
	"errors"
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
			{Name: "Knit", WeeklyFrequency: 3, Ticks: 2, ID: "ID01"},
			{Name: "Water the plants", WeeklyFrequency: 1, Ticks: 0, ID: "ID02"},
		}, nil
	})

	s := New(cli, t)
	s.index(rr, httptest.NewRequest(http.MethodGet, "/?week=1713045600", nil))

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	expect, err := os.ReadFile("testdata/index.html")
	require.NoError(t, err)
	assert.Equal(t, string(expect), rr.Body.String())
}

func TestServer_index_emptylist(t *testing.T) {
	rr := httptest.NewRecorder()

	cli := mocks.NewHabitsClientMock(t)
	cli.ListHabitsMock.Set(func(_ context.Context, _ time.Time) ([]habit.Habit, error) {
		return []habit.Habit{}, nil
	})

	s := New(cli, t)
	s.index(rr, httptest.NewRequest(http.MethodGet, "/?week=1713045600", nil))

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	expect, err := os.ReadFile("testdata/index_emptylist.html")
	require.NoError(t, err)
	assert.Equal(t, string(expect), rr.Body.String())
}

func TestServer_index_failingClient(t *testing.T) {
	testCases := map[string]struct {
		clientErr  error
		wantStatus int
	}{
		"500": {
			clientErr:  errors.New("heute leider nicht"),
			wantStatus: http.StatusInternalServerError,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {

			rr := httptest.NewRecorder()

			cli := mocks.NewHabitsClientMock(t)
			cli.ListHabitsMock.Set(func(_ context.Context, _ time.Time) ([]habit.Habit, error) {
				return nil, testCase.clientErr
			})

			s := New(cli, t)
			s.index(rr, httptest.NewRequest(http.MethodGet, "/", nil))

			assert.Equal(t, testCase.wantStatus, rr.Result().StatusCode)
		})
	}
}
