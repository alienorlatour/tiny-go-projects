package getstatus

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/repository"
)

func TestHandler(t *testing.T) {
	tt := map[string]struct {
		wantStatusCode int
		wantBody       string
		finder         gameFinder
	}{
		"nominal in progress": {
			wantStatusCode: http.StatusOK,
			wantBody:       `{"id":"123456","attempts_left":3,"guesses":[],"word_length":0,"status":"Playing"}`,
			finder: gameFinderStub{
				game: domain.Game{
					ID:           "123456",
					AttemptsLeft: 3,
					Guesses:      nil,
					Status:       domain.StatusPlaying,
				},
				err: nil,
			},
		},
		"nominal won": {
			wantStatusCode: http.StatusOK,
			wantBody:       `{"id":"123456","attempts_left":3,"guesses":[],"word_length":0,"status":"Won"}`,
			finder: gameFinderStub{
				game: domain.Game{
					ID:           "123456",
					AttemptsLeft: 3,
					Guesses:      nil,
					Status:       domain.StatusWon,
				},
				err: nil,
			},
		},
		"not found": {
			wantStatusCode: http.StatusNotFound,
			finder:         gameFinderStub{err: repository.ErrNotFound},
		},
		"other error": {
			wantStatusCode: http.StatusInternalServerError,
			finder:         gameFinderStub{err: fmt.Errorf("not today")},
		},
	}

	for name, testCase := range tt {

		t.Run(name, func(t *testing.T) {
			f := Handler(testCase.finder)

			// Create a request to pass to our handler.
			path := strings.Replace(api.GetStatusPath, fmt.Sprintf("{%s}", api.GameID), "123456", 1)
			req, err := http.NewRequest(http.MethodGet, path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// add path parameters
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add(api.GameID, "123456")
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			f.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			assert.Equal(t, testCase.wantStatusCode, rr.Code)

			// Check the response body is what we expect. Use JSONEq rather than Equal.
			if testCase.wantBody != "" {
				assert.JSONEq(t, testCase.wantBody, rr.Body.String())
			}
		})

	}
}

func TestHandler_missingParameter(t *testing.T) {
	f := Handler(nil)

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	req, err := http.NewRequest(http.MethodGet, "/games", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	f.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	assert.Equal(t, http.StatusNotFound, rr.Code)
}

type gameFinderStub struct {
	game domain.Game
	err  error
}

func (g gameFinderStub) Find(_ domain.GameID) (domain.Game, error) {
	return g.game, g.err
}
