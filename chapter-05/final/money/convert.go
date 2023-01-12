package money

import (
	"context"
	"fmt"
)

const (
	// ErrUnknownChangeRate is returned if we can't find the conversion rate between two currencies.
	ErrUnknownChangeRate = moneyError("no change rate known between currencies")
)

type rateRepository interface {
	// ExchangeRate fetches the ChangeRate for the day and returns it.
	ExchangeRate(ctx context.Context, source, target Currency) (ChangeRate, error)
}

// Convert applies the change rate to convert the given amount to the target currency.
func Convert(ctx context.Context, number Number, from, to Currency, rateRepo rateRepository) (string, error) {
	// validate the given amount is in the handled bounded range
	if err := number.validateInput(from); err != nil {
		return "", err
	}

	// fetch the change rate for the day
	r, err := fetchChangeRate(ctx, from, to, rateRepo)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrUnknownChangeRate, err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := number.applyChangeRate(r, 2)

	// validate the converted amount is in the handled bounded range
	if err := convertedValue.validateOutput(to); err != nil {
		return "", err
	}

	// format the converted value to a readable format
	return convertedValue.String(), nil
}

// fetchChangeRate is in charge of retrieving the change rate between two currencies.
func fetchChangeRate(ctx context.Context, from, to Currency, rateRepo rateRepository) (ChangeRate, error) {
	exchangeRate, err := rateRepo.ExchangeRate(ctx, from, to)
	if err != nil {
		return 0, fmt.Errorf("unable to get exchange rates: %w", err)
	}

	return exchangeRate, nil
}
