package ecbank

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
)

const (
	ErrBadURL             = ecbankError("unable to build http request")
	ErrServerSide         = ecbankError("error from server")
	ErrChangeRateNotFound = ecbankError("couldn't find the exchange rate")
)

type ExchangeRate struct {
	exchangeRatesURL string // "https://www.ecb.europa.eu"
}

func New(url string) *ExchangeRate {
	return &ExchangeRate{url}
}

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (er ExchangeRate) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {
	// add a timeout to the context in case the external API is too slow
	getCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// build the HTTP request
	path, err := er.url()
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrBadURL, err)
	}

	req, err := http.NewRequestWithContext(getCtx, http.MethodGet, path, nil)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrBadURL, path)
	}

	// use the default http client provided by the http library
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("http error:", err)
		return 0., fmt.Errorf("%w: %s", ErrServerSide, path)
	}

	// don't forget to close the response's body
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		// not exposing the API error, should be logged for debug purposes
		return 0., err
	}

	// read the response
	decoder := xml.NewDecoder(resp.Body)

	var ecbMessage Envelope
	err = decoder.Decode(&ecbMessage)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrServerSide, err)
	}

	rate, err := ecbMessage.changeRate(source, target)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}

	return rate, nil
}

const (
	Host string = "https://www.ecb.europa.eu/"
	base string = "/stats/eurofxref/eurofxref-daily.xml"
)

func (er ExchangeRate) url() (string, error) {
	return url.JoinPath(er.exchangeRatesURL, base)
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
		return ecbankError(fmt.Sprintf("user-end error: %d", statusCode))
	case statusCode/httpErrorCategorySize == serverErrorHundreds:
		// errors 5xx
		return ecbankError(fmt.Sprintf("server error: %d", statusCode))
	case statusCode != http.StatusOK:
		// any other usecases
		return ecbankError(fmt.Sprintf("invalid status code: %d", statusCode))
	}
	return nil
}
