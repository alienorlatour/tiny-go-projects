package api

import "learngo-pockets/httpgordle/internal/session"

// ToGameResponse converts a domain.Game into an api.Response.
func ToGameResponse(g session.Game) GameResponse {
	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		Status:       string(g.Status),
		// TODO WordLength
	}

	for index := 0; index < len(g.Guesses); index++ {
		apiGame.Guesses[index].Word = g.Guesses[index].Word
		apiGame.Guesses[index].Feedback = g.Guesses[index].Feedback
	}

	if g.AttemptsLeft == 0 {
		apiGame.Solution = "" // missing solution
	}

	return apiGame
}
