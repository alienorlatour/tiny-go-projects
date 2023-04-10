package ecbank

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"learngo-pockets/moneyconverter/money"
)

const (
	ErrCallingServer      = ecbankError("error calling server")
	ErrTimeout            = ecbankError("timed out when waiting for response")
	ErrUnexpectedFormat   = ecbankError("unexpected response format")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
	ErrClientSide         = ecbankError("client side error when contacting ECB")
	ErrServerSide         = ecbankError("server side error when contacting ECB")
	ErrUnknownStatusCode  = ecbankError("unknown status code contacting ECB")
)

// Client can call the bank to retrieve exchange rates.
type Client struct {
	client *http.Client
}

func NewBank(timeout time.Duration) Client {
	return Client{
		client: &http.Client{Timeout: timeout},
	}
}

// FetchExchangeRate fetches today's ExchangeRate and returns it.
func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	const euroxrefURL = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := c.client.Get(euroxrefURL)
	if err != nil {
		var urlErr *url.Error
		if ok := errors.As(err, &urlErr); ok && urlErr.Timeout() {
			return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrTimeout, err.Error())
		}
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
	userErrorClass     = 4
	serverErrorClass   = 5
	httpErrorClassSize = 100
)

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {
	case statusCode == http.StatusOK:
		return nil
	case statusCode/httpErrorClassSize == userErrorClass:
		// errors 4xx
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case statusCode/httpErrorClassSize == serverErrorClass:
		// errors 5xx
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		// any other usecases
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}
