package apiconversion

import (
	"testing"

	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
	"learngo-pockets/httpgordle/internal/gordle"

	"github.com/stretchr/testify/assert"
)

func TestToAPIResponse(t *testing.T) {
	id := "1682279480"
	tt := map[string]struct {
		game domain.Game
		want api.GameResponse
	}{
		"nominal": {
			game: domain.Game{
				ID: domain.GameID(id),
				Gordle: func() gordle.Game {
					g, _ := gordle.New([]string{"HELLO"})
					return *g
				}(),
				AttemptsLeft: 4,
				Guesses: []domain.Guess{{
					Word:     "FAUNE",
					Feedback: "⬜️🟡⬜️⬜️⬜️",
				}},
				Status: domain.StatusPlaying,
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
				Status:     domain.StatusPlaying,
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
