package money

import (
	"context"
	"fmt"
)

const (
	// ErrGettingChangeRate is returned if we can't find the conversion rate between two currencies.
	ErrGettingChangeRate = moneyError("can't get change rate between currencies")
)

type exchangeRates interface {
	// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
	FetchExchangeRate(ctx context.Context, source, target Currency) (ExchangeRate, error)
}

// Convert parses the input amount and applies the change rate to convert it to the target currency.
func Convert(ctx context.Context, amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	// validate the given amount is in the handled bounded range
	if err := amount.Number.validateInput(amount.Currency); err != nil {
		return Amount{}, err
	}

	// fetch the change rate for the day
	r, err := rates.FetchExchangeRate(ctx, amount.Currency, to)
	if err != nil {
		return Amount{}, fmt.Errorf("%w: %s", ErrGettingChangeRate, err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := amount.applyChangeRate(r, to)

	// validate the converted amount is in the handled bounded range
	if err := convertedValue.Number.validateOutput(to); err != nil {
		return Amount{}, err
	}

	// transform the converted value to an amount
	return convertedValue, nil
}
