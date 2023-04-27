package api

// GuessRequest is the structure of the message used when submitting a guess.
type GuessRequest struct {
	Guess string `json:"guess"`
}
