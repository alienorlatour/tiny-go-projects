package ecbank

// ecbankError defines an error.
type ecbankError string

// ecbankError implements the error interface.
func (e ecbankError) Error() string {
	return string(e)
}
