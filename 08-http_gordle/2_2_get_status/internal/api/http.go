package api

const (
	// GameID is the name of the field that stores the game's identifier
	GameID = "id"

	// NewGameRoute is the path to create a new game.
	NewGameRoute = "/games"
	// GetStatusRoute is the path to get the status of a game identified by its id.
	GetStatusRoute = "/games/{" + GameID + "}"
	// GuessRoute is the path to play a guess in a game, identified by its id.
	GuessRoute = "/games/{" + GameID + "}"
)

// GuessRequest is the structure of the message used when submitting a guess.
type GuessRequest struct {
	// Guess is a word attempted by the player.
	Guess string `json:"guess"`
}

// GameResponse contains the information about a game.
type GameResponse struct {
	// ID is the identified of a game.
	ID string `json:"id"`
	// AttemptsLeft counts the number of attempts left before the game is over.
	AttemptsLeft byte `json:"attempts_left"`
	// Guesses is the list of past guesses, and their feedback.
	Guesses []Guess `json:"guesses"`
	// WordLength is the number of characters in the word to guess.
	WordLength byte `json:"word_length"`
	// Solution is Gordle's secret word. It is only provided when there are no attempts left.
	Solution string `json:"solution,omitempty"`
	// Status displays whether the game is still playable.
	Status string `json:"status"`
}

// A Guess is a pair of a word (submitted by the player) and its feedback (provided by Gordle).
type Guess struct {
	Word     string `json:"word"`
	Feedback string `json:"feedback"`
}
