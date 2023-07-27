package api

import "learngo-pockets/httpgordle/internal/session"

// ToGameResponse converts a domain.Game into an api.Response.
func ToGameResponse(g session.Game) GameResponse {
	solution := g.Gordle.ShowAnswer()

	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		Status:       string(g.Status),
		WordLength:   byte(len(solution)),
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
