package repository

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

type ChangeRateRepository struct {
	exchangeRatesURL string // "https://www.ecb.europa.eu"
}

func New(url string) *ChangeRateRepository {
	return &ChangeRateRepository{url}
}

// ExchangeRate fetches the ChangeRate for the day and returns it.
func (crr ChangeRateRepository) ExchangeRate(source, target money.Currency) (money.ChangeRate, error) {
	// build the HTTP request
	req, err := http.NewRequest(http.MethodGet, crr.url(), nil)
	if err != nil {
		return 0., fmt.Errorf("unable to build http request to %s: %w", crr.url(), err)
	}

	// use the default http client provided by the http library
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0., fmt.Errorf("error while requesting URL %s: %w", crr.url(), err)
	}

	// don't forget to close the response's body
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		// TODO: have our own error?
		return 0., fmt.Errorf("invalid reponse status code: %w", err)
	}

	// read the response
	decoder := xml.NewDecoder(resp.Body)

	var ecbMessage envelope
	err = decoder.Decode(&ecbMessage)
	if err != nil {
		return 0., fmt.Errorf("unable to decode message: %w", err)
	}

	// do we want to return his directly ?
	rate, err := ecbMessage.changeRate(source, target)
	if err != nil {
		return 0., fmt.Errorf("couldn't find the exchange rate: %w", err)
	}

	return rate, nil
}

const (
	ECBRepoURL  = "https://www.ecb.europa.eu/"
	euroxfRoute = "/stats/eurofxref/eurofxref-daily.xml"
)

func (crr ChangeRateRepository) url() string {
	return crr.exchangeRatesURL + euroxfRoute
}
