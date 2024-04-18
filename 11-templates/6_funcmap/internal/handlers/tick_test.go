package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"learngo-pockets/templates/internal/habit"
	"learngo-pockets/templates/internal/handlers/mocks"

	"github.com/stretchr/testify/assert"
)

func TestServer_Tick(t *testing.T) {
	rr := httptest.NewRecorder()

	cli := mocks.NewHabitsClientMock(t)
	cli.TickHabitMock.Set(func(_ context.Context, _ habit.ID) error {
		return nil
	})

	s := New(cli, t)
	s.tick(rr, httptest.NewRequest(http.MethodGet, "/", nil))

	assert.Equal(t, http.StatusSeeOther, rr.Result().StatusCode)
	assert.Contains(t, rr.Body.String(), `<a href="/">`)
}

func TestServer_Tick_error(t *testing.T) {
	rr := httptest.NewRecorder()
	sentinelErr := errors.New("heute leider nicht")

	cli := mocks.NewHabitsClientMock(t)
	cli.TickHabitMock.Set(func(_ context.Context, _ habit.ID) error {
		return sentinelErr
	})

	s := New(cli, t)
	s.tick(rr, httptest.NewRequest(http.MethodGet, "/", nil))

	assert.Equal(t, http.StatusInternalServerError, rr.Result().StatusCode)
}
