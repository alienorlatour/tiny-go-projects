package money

// moneyError defines a sentinel error.
type moneyError string

// moneyError implements the error interface.
func (e moneyError) Error() string {
	return string(e)
}
