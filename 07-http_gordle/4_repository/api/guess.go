package api

// GuessRequest is the structure of the message used when submitting a guess.
type GuessRequest struct {
	// Guess is a word attempted by the player.
	Guess string `json:"guess"`
}
