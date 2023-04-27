package apiconversion

import (
	"learngo-pockets/httpgordle/api"
	"learngo-pockets/httpgordle/internal/domain"
)

func ToAPIResponse(g domain.Game) api.GameResponse {
	apiGame := api.GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]api.Guess, len(g.Guesses)),
		WordLength:   byte(len(g.Solution)),
	}

	for index := 0; index < len(g.Guesses); index++ {
		apiGame.Guesses[index].Word = g.Guesses[index].Word
		apiGame.Guesses[index].Feedback = g.Guesses[index].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = g.Solution
	}

	return apiGame
}
