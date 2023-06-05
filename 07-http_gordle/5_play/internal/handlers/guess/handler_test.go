package guess

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/session"
)

func TestHandle(t *testing.T) {
	game, _ := gordle.New([]string{"pocket"})
	handle := Handler(successGameGuesserStub{session.Game{
		ID:           "123456",
		Gordle:       *game,
		AttemptsLeft: 5,
		Guesses:      nil,
		Status:       session.StatusPlaying,
	}})

	req, err := http.NewRequest(http.MethodPost, "/games/123456", strings.NewReader(`{"guess":"pocket"}`))
	require.NoError(t, err)

	// add path parameters
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add(api.GameID, "123456")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	recorder := httptest.NewRecorder()

	handle(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, `{"id":"123456","attempts_left":4,"guesses":[{"word":"pocket", "feedback":"++++++"}],"word_length":6,"status":"Won"}`, recorder.Body.String())
}

type successGameGuesserStub struct {
	game session.Game
}

func (g successGameGuesserStub) Find(id session.GameID) (session.Game, error) {
	return g.game, nil
}

func (g successGameGuesserStub) Update(game session.Game) error {
	return nil
}
