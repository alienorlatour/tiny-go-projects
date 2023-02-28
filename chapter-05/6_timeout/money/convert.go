package money

import (
	"fmt"
)

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	// fetch the change rate for the day
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get change rate: %w", err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue, err := applyChangeRate(amount, to, r)
	if err != nil {
		return Amount{}, err
	}

	// validate the converted amount is in the handled bounded range
	if err = convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	return convertedValue, nil
}

type exchangeRates interface {
	// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}

// ExchangeRate represents a rate to convert from a currency to another.
// It is a float64, because the precision of an official change rate is 5 significant digits.
type ExchangeRate float64

// applyChangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyChangeRate(a Amount, target Currency, rate ExchangeRate) (Amount, error) {
	// Multiply the input amount.
	converted, err := multiply(a.quantity, rate)
	if err != nil {
		return Amount{}, err
	}

	// Adjust precision
	switch {
	case converted.precision > target.precision:
		// The converted value is too precise, let's chunk some digits off. This will floor down the result.
		converted.subunits = converted.subunits / tenToThe(converted.precision-target.precision)
		converted.precision = target.precision
	case converted.precision < target.precision:
		// Multiply, adding enough zeroes to reach the desired precision.
		converted.subunits = converted.subunits * tenToThe(target.precision-converted.precision)
		converted.precision = target.precision
	}

	return Amount{
		currency: target,
		quantity: converted,
	}, nil
}

func multiply(d Decimal, r ExchangeRate) (Decimal, error) {
	// first, convert the ExchangeRate to a Decimal
	rate, err := ParseDecimal(fmt.Sprintf("%f", r))
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: exchange rate is %f", ErrInvalidDecimal, r)
	}

	return Decimal{
		subunits:  d.subunits * rate.subunits,
		precision: d.precision + rate.precision,
	}, nil
}
