// Package ecbank exposes a way to query the API of the European Central Bank.
package ecbank

import (
	"fmt"
	"net/http"
	"time"

	"learngo-pockets/moneyconverter/money"
)

const (
	ErrServerSide         = Error("error from server")
	ErrUnexpectedFormat   = Error("unexpected response format")
	ErrChangeRateNotFound = Error("couldn't find the exchange rate")
)

// EuroCentralBank can call the bank to retrieve exchange rates.
type EuroCentralBank struct {
	client http.Client
}

// NewBank builds a EuroCentralBank that can fetch exchange rates within a given timeout.
func NewBank(timeout time.Duration) EuroCentralBank {
	return EuroCentralBank{
		client: http.Client{Timeout: timeout},
	}
}

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (ecb EuroCentralBank) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const path = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := ecb.client.Get(path)
	if err != nil {
		return money.ExchangeRate(0.), fmt.Errorf("%w: %s", ErrServerSide, err.Error())
	}

	// don't forget to close the response's body
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return 0., err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return 0., err
	}

	return rate, nil
}

const (
	userErrorHundreds     = 4
	serverErrorHundreds   = 5
	httpErrorCategorySize = 100
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode/httpErrorCategorySize == userErrorHundreds:
		// errors 4xx
		return Error(fmt.Sprintf("user-end error: %d", statusCode))
	case statusCode/httpErrorCategorySize == serverErrorHundreds:
		// errors 5xx
		return Error(fmt.Sprintf("server error: %d", statusCode))
	case statusCode != http.StatusOK:
		// any other usecases
		return Error(fmt.Sprintf("invalid status code: %d", statusCode))
	}
	return nil
}
