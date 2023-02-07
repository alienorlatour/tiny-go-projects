package money

import (
	"fmt"
	"math"
)

const (
	// ErrGettingChangeRate is returned if we can't find the conversion rate between two currencies.
	ErrGettingChangeRate = moneyError("can't get change rate between currencies")
)

// Convert parses the input amount and applies the change rate to convert it to the target currency.
func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	// fetch the change rate for the day
	r, err := rates.FetchExchangeRate(amount.currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("%w: %s", ErrGettingChangeRate, err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := applyChangeRate(amount, r, to)

	// validate the converted amount is in the handled bounded range
	if err := convertedValue.validate(); err != nil {
		return Amount{}, err
	}

	// transform the converted value to an amount
	return convertedValue, nil
}

// applyChangeRate returns a new Amount representing a multiplied by the rate.
// This function does not guarantee that the output amount is supported by the rest of the library.
func applyChangeRate(a Amount, rate ExchangeRate, target Currency) Amount {
	converted := a.number.float() * float64(rate)

	floor := math.Floor(converted)
	decimal := math.Round((converted - floor) * math.Pow10(target.precision))

	amount := Amount{
		number: Number{
			integerPart: int(floor),
			decimalPart: int(decimal),
			precision:   target.precision,
		},
		currency: target,
	}

	return amount
}
