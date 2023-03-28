package money

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency) (Amount, error) {
	// Convert to the target currency applying the fetched change rate.
	convertedValue := applyExchangeRate(amount, to, ExchangeRate{subunits: 2, precision: 0})

	// validate the converted amount is in the handled bounded range.
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

// ExchangeRate represents a rate to convert from a currency to another.
// It is a Decimal, to support various precisions.
type ExchangeRate Decimal

// applyExchangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {
	// Multiply the input amount.
	converted := multiply(a.quantity, rate)

	// Adjust precision
	switch {
	case converted.precision > target.precision:
		// The converted value is too precise, let's chunk some digits off. This will floor down the result.
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		// Multiply, adding enough zeroes to reach the desired precision.
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)
	}

	converted.precision = target.precision

	return Amount{
		currency: target,
		quantity: converted,
	}
}

// multiply a Decimal with an ExchangeRate and return the product
func multiply(d Decimal, r ExchangeRate) Decimal {
	return Decimal{
		subunits:  d.subunits * r.subunits,
		precision: d.precision + r.precision,
	}
}
