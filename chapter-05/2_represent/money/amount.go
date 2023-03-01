package money

// Amount defines a decimal of money in a given currency.
type Amount struct {
	quantity Decimal
	currency Currency
}

const (
	// ErrTooPrecise is returned if the number is too precise for the currency.
	ErrTooPrecise = Error("quantity is too precise")
)

// NewAmount returns an Amount of money.
func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		// In order to avoid converting 0.00001 cent, let's exit now.
		return Amount{}, ErrTooPrecise
	}

	return Amount{quantity: quantity, currency: currency}, nil
}
