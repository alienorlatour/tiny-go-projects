package session

// domainError is a sentinel error for the session.
type domainError string

// Error implements the error interface.
func (d domainError) Error() string {
	return string(d)
}

const (
	// ErrGameOver is returned when a play is made but the game is over.
	ErrGameOver = domainError("❌game over")
)
