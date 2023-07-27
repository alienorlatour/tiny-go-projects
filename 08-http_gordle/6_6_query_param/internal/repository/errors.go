package repository

// Error is used to define sentinel errors.
type Error string

// Error implements the error interface.
func (r Error) Error() string {
	return string(r)
}

const (
	// ErrNotFound is returned when a game doesn't exist in the repository.
	ErrNotFound = Error("game not found in repository")
	// ErrConflictingID is returned when a game would be created with the same ID as an existing game.
	ErrConflictingID = Error("cannot create game with already-existing ID")
)
