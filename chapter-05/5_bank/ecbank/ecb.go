// Package ecbank exposes a way to query the API of the European Central Bank.
package ecbank

import (
	"fmt"
	"net/http"

	"learngo-pockets/moneyconverter/money"
)

const (
	ErrCallingServer        = ecbankError("error calling server")
	ErrUnexpectedFormat     = ecbankError("unexpected response format")
	ErrChangeRateNotFound   = ecbankError("couldn't find the exchange rate")
	ErrECBUserEnd           = ecbankError("user-end error when contacting ECB")
	ErrECBServerEnd         = ecbankError("server-end error when contacting ECB")
	ErrInvalidECBStatusCode = ecbankError("invalid status code contacting ECB")
)

// EuroCentralBank can call the bank to retrieve exchange rates.
type EuroCentralBank struct {
	path string
}

// FetchExchangeRate fetches today's ExchangeRate and returns it.
func (ecb EuroCentralBank) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const path = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	if ecb.path == "" {
		ecb.path = path
	}

	resp, err := http.Get(ecb.path)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrCallingServer, err.Error())
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
	userErrorHundreds   = 4
	serverErrorHundreds = 5
	httpErrorClassSize  = 100
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case statusCode/httpErrorClassSize == userErrorHundreds:
		// errors 4xx
		return fmt.Errorf("%w: %d", ErrECBUserEnd, statusCode)
	case statusCode/httpErrorClassSize == serverErrorHundreds:
		// errors 5xx
		return fmt.Errorf("%w: %d", ErrECBServerEnd, statusCode)
	default:
		// any other usecases
		return fmt.Errorf("%w: %d", ErrInvalidECBStatusCode, statusCode)
	}
}
