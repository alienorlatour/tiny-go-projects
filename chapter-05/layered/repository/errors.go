package repository

// repositoryError defines a sentinel error.
type repositoryError string

// repositoryError implements the error interface.
func (e repositoryError) Error() string {
	return string(e)
}
