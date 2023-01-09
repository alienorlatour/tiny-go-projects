package money

import (
	"context"
	"fmt"
)

const (
	// ErrUnknownChangeRate is returned if we can't find the conversion rate between two currencies.
	ErrUnknownChangeRate = moneyError("no change rate known between currencies")
	// ErrInputTooSmall is returned if the amount to convert is too small, which could lead to precision issues.
	ErrInputTooSmall = moneyError("input amount should be at least 1.00")
	// ErrOutputTooSmall is returned if the converted amount is too small, in order to protect against precision issues.
	ErrOutputTooSmall = moneyError("output amount is less than 1.00")
	// ErrInputTooLarge is returned if the input amount is too large - this would cause floating point precision errors.
	ErrInputTooLarge = moneyError("input amount over 10^15 is too large")
	// ErrOutputTooLarge is returned if the converted amount is too large, to protect against floating point errors.
	ErrOutputTooLarge = moneyError("output amount is too large (over 10^15)")
)

type rateRepository interface {
	// ExchangeRate fetches the ChangeRate for the day and returns it.
	ExchangeRate(ctx context.Context, source, target Currency) (ChangeRate, error)
}

// Convert parses the input amount and applies the change rate to convert it to the target currency.
func Convert(ctx context.Context, amount, from, to string, rateRepo rateRepository) (string, error) {
	// parse the amount to convert
	n, err := parseNumber(amount)
	if err != nil {
		return "", fmt.Errorf("unable to parse amount: %w", err)
	}

	// validateInput the given amount is in the handled bounded range
	if err := n.validateInput(); err != nil {
		return "", err
	}

	// fetch the change rate for the day
	r, err := fetchChangeRate(ctx, from, to, rateRepo)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrUnknownChangeRate, err)
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := n.applyChangeRate(r, 2)

	if err := convertedValue.validateOutput(); err != nil {
		return "", err
	}

	// format the converted value to a readable format
	return convertedValue.String(), nil
}

func (n number) validateInput() error {
	switch {
	case n.tooSmall():
		return ErrInputTooSmall
	case n.tooBig():
		return ErrInputTooLarge
	}
	return nil
}

func (n number) validateOutput() error {
	switch {
	case n.tooSmall():
		return ErrOutputTooSmall
	case n.tooBig():
		return ErrOutputTooLarge
	}
	return nil
}

// fetchChangeRate is in charge of retrieving the change rate between two currencies.
func fetchChangeRate(ctx context.Context, from, to string, rateRepo rateRepository) (ChangeRate, error) {
	// get the output currency
	sourceCurrency, err := parseCurrency(from)
	if err != nil {
		return 0, fmt.Errorf("unable to parse source currency: %w", err)
	}

	targetCurrency, err := parseCurrency(to)
	if err != nil {
		return 0, fmt.Errorf("unable to parse target currency: %w", err)
	}

	exchangeRate, err := rateRepo.ExchangeRate(ctx, sourceCurrency, targetCurrency)
	if err != nil {
		return 0, fmt.Errorf("unable to get exchange rates: %w", err)
	}

	return exchangeRate, nil
}
