package money

import (
	"fmt"
)

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency, rates ratesFetcher) (Amount, error) {
	// fetch the change rate for the day
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get change rate: %w", err)
	}

	// Convert to the target currency applying the fetched change rate.
	convertedValue, err := applyExchangeRate(amount, to, r)
	if err != nil {
		return Amount{}, err
	}

	// Validate the converted amount is in the handled bounded range.
	if err = convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

type ratesFetcher interface {
	// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}

// ExchangeRate represents a rate to convert from a currency to another.
// It is a Decimal, to support various precisions.
type ExchangeRate Decimal

// applyExchangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) (Amount, error) {
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
	}, nil
}

// multiply a Decimal with an ExchangeRate and return the product
func multiply(d Decimal, r ExchangeRate) Decimal {
	dec := Decimal{
		subunits:  d.subunits * r.subunits,
		precision: d.precision + r.precision,
	}

	// Clean the product before returning it.
	dec.simplify()

	return dec
}
