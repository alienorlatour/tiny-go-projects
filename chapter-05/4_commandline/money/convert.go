package money

import "math"

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	// convert to the target currency applying the fetched change rate
	convertedValue := applyChangeRate(amount, to, 2)

	// validate the converted amount is in the handled bounded range
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

// ExchangeRate represents a rate to convert from a currency to another.
// It is a float64, because the precision of an official change rate is 5 significant digits.
type ExchangeRate float64

// applyChangeRate returns a new Number representing n multiplied by the rate.
// The precision is the same in and out.
// This function does not guarantee that the output amount is supported.
func applyChangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	amount := Amount{
		currency: target,
		quantity: Quantity{
			exp: target.precision,
		},
	}

	cents := float64(a.quantity.cents) * float64(rate) * math.Pow10(target.precision-a.quantity.exp)

	amount.quantity.cents = int(math.Floor(cents))
	return amount
}
