package gordle

// corpusError defines a sentinel error.
type corpusError string

// Error is the implementation of the error interface by corpusError.
func (e corpusError) Error() string {
	return string(e)
}

// gameError defines an error that happens during a game.
type gameError string

// Error is the implementation of the error interface by gameError
func (e gameError) Error() string {
	return string(e)
}
