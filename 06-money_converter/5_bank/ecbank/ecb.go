// Package ecbank exposes a way to query the API of the European Central Bank.
package ecbank

import (
	"fmt"
	"net/http"

	"learngo-pockets/moneyconverter/money"
)

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrUnexpectedFormat   = ecbankError("unexpected response format")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
	ErrClientSide         = ecbankError("client side error when contacting ECB")
	ErrServerSide         = ecbankError("server side error when contacting ECB")
	ErrUnknownStatusCode  = ecbankError("unknown status code contacting ECB")
)

// Client can call the bank to retrieve exchange rates.
type Client struct {
	url string
}

// FetchExchangeRate fetches today's ExchangeRate and returns it.
func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const euroxrefURL = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	if c.url == "" {
		c.url = euroxrefURL
	}

	resp, err := http.Get(c.url)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrCallingServer, err.Error())
	}

	// don't forget to close the response's body
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}

	return rate, nil
}

const (
	clientErrorClass = 4
	serverErrorClass = 5
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		// errors 4xx
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		// errors 5xx
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		// any other usecases
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

// httpStatusClass returns the class of a http status code.
func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}
