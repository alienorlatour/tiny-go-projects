package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"learngo-pockets/templates/internal/habit"
	"learngo-pockets/templates/internal/handlers/mocks"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestServer_Tick(t *testing.T) {
	rr := httptest.NewRecorder()

	req := httptest.NewRequest(http.MethodGet, "/tick/{habitID}", nil)
	rCtx := chi.NewRouteContext()
	rCtx.URLParams.Add("habitID", "1234")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rCtx))

	cli := mocks.NewHabitsClientMock(t)
	cli.TickHabitMock.Expect(req.Context(), "1234").Return(nil)

	s := New(cli, t)

	s.tick(rr, req)

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
