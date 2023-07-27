package newgame

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"learngo-pockets/httpgordle/internal/session"
)

func TestHandle(t *testing.T) {
	handleFunc := Handler(gameAdderStub{})

	req, err := http.NewRequest(http.MethodPost, "/games", nil)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	handleFunc(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())
}

type gameAdderStub struct {
	err error
}

func (g gameAdderStub) Add(_ session.Game) error {
	return g.err
}
