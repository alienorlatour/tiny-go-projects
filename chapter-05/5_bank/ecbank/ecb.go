// Package ecbank exposes a way to query the API of the European Central Bank.
package ecbank

import (
	"fmt"
	"net/http"
	"time"

	"learngo-pockets/moneyconverter/money"
)

const (
	ErrServerSide           = ecbankError("error from server")
	ErrUnexpectedFormat     = ecbankError("unexpected response format")
	ErrChangeRateNotFound   = ecbankError("couldn't find the exchange rate")
	ErrECBUserEnd           = ecbankError("user-end error when contacting ECB")
	ErrECBServerEnd         = ecbankError("server-end error when contacting ECB")
	ErrInvalidECBStatusCode = ecbankError("invalid status code contacting ECB")
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

// FetchExchangeRate fetches today's ExchangeRate and returns it.
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
