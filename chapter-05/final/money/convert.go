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
func Convert(ctx context.Context, amount Number, from, to Currency, rates exchangeRates) (string, error) {
	// validate the given amount is in the handled bounded range
	if err := amount.validateInput(from); err != nil {
		return "", err
	}

	// fetch the change rate for the day
	r, err := fetchExchangeRate(ctx, from, to, rates)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrGettingChangeRate, err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := amount.applyChangeRate(r, 2)

	// validate the converted amount is in the handled bounded range
	if err := convertedValue.validateOutput(to); err != nil {
		return "", err
	}

	// format the converted value to a readable format
	return convertedValue.String(), nil
}

// fetchExchangeRate is in charge of retrieving the change rate between two currencies.
func fetchExchangeRate(ctx context.Context, from, to Currency,  rateRepo exchangeRates) (ExchangeRate, error) {
	exchangeRate, err := rateRepo.FetchExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("unable to get exchange rates: %w", err)
	}

	return exchangeRate, nil
}
