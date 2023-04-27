package api

const (
	// GameID is the name of the field that stores the game's identifier
	GameID = "id"

	// NewGamePath is the path to create a new game.
	NewGamePath = "/games"
	// GetStatusPath is the path to get the status of a game identified by its id.
	GetStatusPath = "/games/{" + GameID + "}"
	// GuessPath is the path to play a guess in a game, identified by its id.
	GuessPath = "/games/{" + GameID + "}"
)
