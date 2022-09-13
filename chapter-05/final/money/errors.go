package money

// internalError defines a sentinel error.
type internalError string

// internalError implements the error interface.
func (e internalError) Error() string {
	return string(e)
}
