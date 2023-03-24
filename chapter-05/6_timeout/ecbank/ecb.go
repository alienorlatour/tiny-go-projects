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
	ErrCallingServer        = ecbankError("error calling server")
	ErrTimeout              = ecbankError("timed out when waiting for response")
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
		var urlErr *url.Error
		if ok := errors.As(err, &urlErr); ok && urlErr.Timeout() {
			return 0., fmt.Errorf("%w: %s", ErrTimeout, err.Error())
		}
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
