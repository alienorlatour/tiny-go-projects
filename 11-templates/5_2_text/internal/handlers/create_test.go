package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"learngo-pockets/templates/internal/habit"
	"learngo-pockets/templates/internal/handlers/mocks"

	"github.com/stretchr/testify/assert"
)

func TestServer_Create(t *testing.T) {
	testCases := map[string]struct {
		input      func() url.Values
		wantStatus int
		client     func() habitsClient
	}{
		"nominal": {
			input: func() url.Values {
				v := url.Values{}
				v.Add("habitName", "Dance your heart out")
				v.Add("habitFrequency", "2")
				return v
			},
			wantStatus: http.StatusSeeOther,
			client: func() habitsClient {
				cli := mocks.NewHabitsClientMock(t)
				cli.CreateHabitMock.Expect(context.Background(), habit.Habit{Name: "Dance your heart out", WeeklyFrequency: 2}).Return(nil)
				return cli
			},
		},
		"client error": {
			input: func() url.Values {
				v := url.Values{}
				v.Add("habitName", "Dance your heart out")
				v.Add("habitFrequency", "2")
				return v
			},
			wantStatus: http.StatusInternalServerError,
			client: func() habitsClient {
				cli := mocks.NewHabitsClientMock(t)
				cli.CreateHabitMock.Expect(context.Background(), habit.Habit{Name: "Dance your heart out", WeeklyFrequency: 2}).
					Return(errors.New("nope"))
				return cli
			},
		},
		"not a number": {
			input: func() url.Values {
				v := url.Values{}
				v.Add("habitName", "Dance your heart out")
				v.Add("habitFrequency", "NaN")
				return v
			},
			wantStatus: http.StatusBadRequest,
			client: func() habitsClient {
				return nil
			},
		},
		"number too high": {
			input: func() url.Values {
				v := url.Values{}
				v.Add("habitName", "Dance your heart out")
				v.Add("habitFrequency", "999")
				return v
			},
			wantStatus: http.StatusBadRequest,
			client: func() habitsClient {
				return nil
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			s := New(testCase.client(), t)
			request := httptest.NewRequest(http.MethodPost, "/create", nil)
			request.Form = testCase.input()

			s.create(rr, request)

			assert.Equal(t, testCase.wantStatus, rr.Result().StatusCode)
		})
	}
}
