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
func Convert(ctx context.Context, amount, from, to string, rates exchangeRates) (string, error) {
	// parse the amount to convert
	n, err := parseNumber(amount)
	if err != nil {
		return "", fmt.Errorf("unable to parse amount: %w", err)
	}

	// validateInput the given amount is in the handled bounded range
	if err = n.validateInput(); err != nil {
		return "", err
	}

	// fetch the change rate for the day
	r, err := fetchExchangeRate(ctx, from, to, rates)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrGettingChangeRate, err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := n.applyChangeRate(r, 2)

	if err := convertedValue.validateOutput(); err != nil {
		return "", err
	}

	// format the converted value to a readable format
	return convertedValue.String(), nil
}

// fetchExchangeRate is in charge of retrieving the change rate between two currencies.
func fetchExchangeRate(ctx context.Context, from, to string, rateRepo exchangeRates) (ExchangeRate, error) {
	// get the output currency
	sourceCurrency, err := parseCurrency(from)
	if err != nil {
		return 0, fmt.Errorf("unable to parse source currency: %w", err)
	}

	targetCurrency, err := parseCurrency(to)
	if err != nil {
		return 0, fmt.Errorf("unable to parse target currency: %w", err)
	}

	exchangeRate, err := rateRepo.FetchExchangeRate(ctx, sourceCurrency, targetCurrency)
	if err != nil {
		return 0, fmt.Errorf("unable to get exchange rates: %w", err)
	}

	return exchangeRate, nil
}
