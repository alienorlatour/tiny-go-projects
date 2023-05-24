package apiconversion

import (
	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/session"
)

func ToAPIResponse(g session.Game) api.GameResponse {
	solution := g.Gordle.ShowAnswer()

	apiGame := api.GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]api.Guess, len(g.Guesses)),
		WordLength:   byte(len(solution)),
		Status:       string(g.Status),
	}

	for i := range apiGame.Guesses {
		apiGame.Guesses[i].Word = g.Guesses[i].Word
		apiGame.Guesses[i].Feedback = g.Guesses[i].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = solution
	}

	return apiGame
}
