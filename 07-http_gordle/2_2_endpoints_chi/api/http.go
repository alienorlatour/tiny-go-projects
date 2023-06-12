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
