package gordle

// corpusError defines a sentinel error.
type corpusError string

// Error is the implementation of the error interface by corpusError.
func (e corpusError) Error() string {
	return string(e)
}
