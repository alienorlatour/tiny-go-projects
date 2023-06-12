package session

// Error is a sentinel error for the domain.
type Error string

// Error implements the error interface.
func (d Error) Error() string {
	return string(d)
}

const (
	// ErrGameOver is returned when a play is made but the game is over.
	ErrGameOver = Error("❌game over")
)
