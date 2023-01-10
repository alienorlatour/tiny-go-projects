package repository

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

const (
	ErrBadURL             = repositoryError("unable to build http request")
	ErrServerSide         = repositoryError("error from server")
	ErrChangeRateNotFound = repositoryError("couldn't find the exchange rate")
)

type ChangeRateRepository struct {
	exchangeRatesURL string // "https://www.ecb.europa.eu"
}

func New(url string) *ChangeRateRepository {
	return &ChangeRateRepository{url}
}

// ExchangeRate fetches the ChangeRate for the day and returns it.
func (crr ChangeRateRepository) ExchangeRate(ctx context.Context, source, target money.Currency) (money.ChangeRate, error) {
	// add a timeout to the context in case the external API is too slow
	getCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// build the HTTP request
	req, err := http.NewRequestWithContext(getCtx, http.MethodGet, crr.url(), nil)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrBadURL, crr.url())
	}

	// use the default http client provided by the http library
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("http error:", err)
		return 0., fmt.Errorf("%w: %s", ErrServerSide, crr.url())
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
	euroxfRoute = "/stats/eurofxref/eurofxref-daily.xml"
)

func (crr ChangeRateRepository) url() string {
	return crr.exchangeRatesURL + euroxfRoute
}
