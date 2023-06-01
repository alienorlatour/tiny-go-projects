package newgame

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/session"
)

func TestHandler(t *testing.T) {
	corpusPath = "testdata/corpus.txt"
	idFinderRegexp := regexp.MustCompile(`.+"id":"(\d+)".+`)

	tt := map[string]struct {
		wantStatusCode int
		wantBody       string
		creator        gameCreator
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

			req, err := http.NewRequest(http.MethodPost, api.NewGamePath, nil)
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
			assert.JSONEq(t, testCase.wantBody, body)
		})
	}
}

type gameCreatorStub struct {
	err error
}

func (g gameCreatorStub) Add(session.Game) error {
	return g.err
}
