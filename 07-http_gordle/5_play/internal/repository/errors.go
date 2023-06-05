package repository

// Error is used to define sentinel errors.
type Error string

// Error implements the error interface.
func (r Error) Error() string {
	return string(r)
}

// ErrNotFound is returned when a game doesn't exist in the repository.
const (
	ErrNotFound      = Error("game not found in repository")
	ErrConflictingID = Error("cannot create game with already-existing ID")
)
