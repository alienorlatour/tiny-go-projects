package repository

// repoErr is used to define sentinel errors.
type repoErr string

// Error implements the error interface.
func (r repoErr) Error() string {
	return string(r)
}

// ErrNotFound is returned when a game doesn't exist in the repository.
const ErrNotFound = repoErr("game not found in repository")
