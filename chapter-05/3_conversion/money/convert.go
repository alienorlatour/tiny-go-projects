package money

import "math"

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	// Convert to the target currency applying the fetched change rate.
	convertedValue := applyChangeRate(amount, to, 2)

	// validate the converted amount is in the handled bounded range.
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

// ExchangeRate represents a rate to convert from a currency to another.
// It is a float64, because the precision of an official change rate is 5 significant digits.
type ExchangeRate float64

// applyChangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyChangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	amount := Amount{
		currency: target,
		quantity: Decimal{
			precision: target.precision,
		},
	}

	// Apply the change rate and use the target's subunit.
	cents := float64(a.quantity.subunits) * float64(rate) * math.Pow10(int(target.precision)-int(a.quantity.precision))

	// We floor the result, which avoids creating money.
	amount.quantity.subunits = int64(math.Floor(cents))
	return amount
}
