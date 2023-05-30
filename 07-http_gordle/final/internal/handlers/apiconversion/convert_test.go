package apiconversion

import (
	"testing"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/gordle"
	"learngo-pockets/httpgordle/internal/session"

	"github.com/stretchr/testify/assert"
)

func TestToAPIResponse(t *testing.T) {
	id := "1682279480"
	tt := map[string]struct {
		game session.Game
		want api.GameResponse
	}{
		"nominal": {
			game: session.Game{
				ID: session.GameID(id),
				Gordle: func() gordle.Game {
					g, _ := gordle.New("HELLO")
					return *g
				}(),
				AttemptsLeft: 4,
				Guesses: []session.Guess{{
					Word:     "FAUNE",
					Feedback: "⬜️🟡⬜️⬜️⬜️",
				}},
				Status: session.StatusPlaying,
			},
			want: api.GameResponse{
				ID:           id,
				AttemptsLeft: 4,
				Guesses: []api.Guess{{
					Word:     "FAUNE",
					Feedback: "⬜️🟡⬜️⬜️⬜️",
				}},
				WordLength: 5,
				Solution:   "",
				Status:     session.StatusPlaying,
			},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := ToAPIResponse(tc.game)
			assert.Equal(t, tc.want, got)
		})
	}
}
