package apiconversion

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/session"
)

func TestToAPIResponse(t *testing.T) {
	id := "1682279480"
	tt := map[string]struct {
		game session.Game
		want api.GameResponse
	}{
		"nominal": {
			game: session.Game{
				ID:           session.GameID(id),
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
				Solution: "",
				Status:   session.StatusPlaying,
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
