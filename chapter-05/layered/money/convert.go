package money

import "fmt"

const (
	errUnknownChangeRate = moneyError("no change rate known between currencies")
)

type rateRepository interface {
	// ExchangeRate fetches the ChangeRate for the day and returns it.
	ExchangeRate(source, target Currency) (ChangeRate, error)
}

// Convert parses the input amount and applies the change rate to convert it to the target currency.
func Convert(amount, from, to string, rateRepo rateRepository) (string, error) {
	// parse the amount to convert
	n, err := parseNumber(amount)
	if err != nil {
		return "", fmt.Errorf("unable to parse amount: %w", err)
	}

	// fetch the change rate for the day
	r, err := fetchChangeRate(from, to, rateRepo)
	if err != nil {
		return "", fmt.Errorf("%w: %s", errUnknownChangeRate, err.Error())
	}

	// convert to the target currency applying the fetched change rate
	convertedValue := n.applyChangeRate(r, 2)

	// format the converted value to a readable format
	return convertedValue.String(), nil
}

// fetchChangeRate is in charge of retrieving the change rate between two currencies.
func fetchChangeRate(from, to string, rateRepo rateRepository) (ChangeRate, error) {
	// get the output currency
	sourceCurrency, err := parseCurrency(from)
	if err != nil {
		return 0, fmt.Errorf("unable to parse source currency: %w", err)
	}

	targetCurrency, err := parseCurrency(to)
	if err != nil {
		return 0, fmt.Errorf("unable to parse target currency: %w", err)
	}

	exchangeRate, err := rateRepo.ExchangeRate(sourceCurrency, targetCurrency)
	if err != nil {
		return 0, fmt.Errorf("unable to get exchange rates: %w", err)
	}

	return exchangeRate, nil
}
