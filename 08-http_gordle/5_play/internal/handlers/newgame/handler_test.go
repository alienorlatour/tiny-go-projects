package newgame

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"learngo-pockets/httpgordle/internal/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/session"
)

func TestHandler(t *testing.T) {
	// idFinderRegexp is a regular expression that will ensure the body contains an id field with a value that contains
	// only letters (uppercase and/or lowercase) and/or digits.
	idFinderRegexp := regexp.MustCompile(`.+"id":"([a-zA-Z0-9]+)".+`)

	tt := map[string]struct {
		wantStatusCode int
		wantBody       string
		creator        gameAdder
	}{
		"nominal": {
			wantStatusCode: http.StatusCreated,
			wantBody:       `{"id":"123456","attempts_left":5,"guesses":[],"word_length":5,"status":"Playing"}`,
			creator: gameCreatorStub{
				err: nil,
			},
		},
	}

	for name, testCase := range tt {

		t.Run(name, func(t *testing.T) {
			f := Handler(testCase.creator)

			req, err := http.NewRequest(http.MethodPost, api.NewGameRoute, nil)
			if err != nil {
				t.Fatal(err)
			}

			// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
			rr := httptest.NewRecorder()

			f.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			assert.Equal(t, testCase.wantStatusCode, rr.Code)

			// Check the response body is what we expect. Use JSONEq rather than Equal.
			if testCase.wantBody == "" {
				return
			}

			// replace the ID
			body := rr.Body.String()
			id := idFinderRegexp.FindStringSubmatch(body)
			if len(id) != 2 {
				t.Fatal("cannot find one single id in the json output")
			}
			body = strings.Replace(body, id[1], "123456", 1)

			// validate the rest
			assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
			assert.JSONEq(t, testCase.wantBody, body)
		})
	}
}

func Test_createGame(t *testing.T) {
	g, err := createGame(gameCreatorStub{nil})
	require.NoError(t, err)

	assert.Equal(t, uint8(5), g.AttemptsLeft)
	assert.Equal(t, 0, len(g.Guesses))
	assert.Regexp(t, "[A-Z0-9]+", g.ID)

	corpus, _ := gordle.ParseCorpus()
	assert.Contains(t, corpus, g.Gordle.ShowAnswer())
}

type gameCreatorStub struct {
	err error
}

func (g gameCreatorStub) Add(session.Game) error {
	return g.err
}
