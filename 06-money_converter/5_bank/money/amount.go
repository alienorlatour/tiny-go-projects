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
	switch {
	case quantity.precision > currency.precision:
		// In order to avoid converting 0.00001 cent, let's exit now.
		return Amount{}, ErrTooPrecise
	case quantity.precision < currency.precision:
		quantity.subunits *= pow10(currency.precision - quantity.precision)
		quantity.precision = currency.precision
	}

	return Amount{quantity: quantity, currency: currency}, nil
}

// validate returns an error if and only if an Amount is unsafe to use.
func (a Amount) validate() error {
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrTooPrecise
	}

	return nil
}

// String implements stringer.
func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.code
}
