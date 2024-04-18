package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"learngo-pockets/templates/internal/habit"
	"learngo-pockets/templates/internal/handlers/mocks"

	"github.com/stretchr/testify/assert"
)

func TestServer_Create(t *testing.T) {
	rr := httptest.NewRecorder()

	cli := mocks.NewHabitsClientMock(t)
	cli.CreateHabitMock.Expect(context.Background(), habit.Habit{Name: "Dance your heart out", WeeklyFrequency: 2}).Return(nil)

	s := New(cli, t)
	request := httptest.NewRequest(http.MethodPost, "/create", nil)
	request.Form = url.Values{}
	request.Form.Add("habitName", "Dance your heart out")
	request.Form.Add("habitFrequency", "2")
	s.create(rr, request)

	assert.Equal(t, http.StatusSeeOther, rr.Result().StatusCode)
}
