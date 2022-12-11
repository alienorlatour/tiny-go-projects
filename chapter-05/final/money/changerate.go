package money

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	errUnknownChangeRate = moneyError("no change rate known between currencies")
)

// changeRate is a float32, because the precision of an official change rate is 5 significant digits.
type changeRate float32

// fetchChangeRate is in charge of retrieving the change rate between two currencies.
func fetchChangeRate(from, to string) (changeRate, error) {
	// get the output currency
	sourceCurrency, err := getCurrency(from)
	if err != nil {
		return 0, fmt.Errorf("unable to parse source currency: %w", err)
	}

	targetCurrency, err := getCurrency(to)
	if err != nil {
		return 0, fmt.Errorf("unable to parse target currency: %w", err)
	}

	exchangeRates, err := getECBExchangeRate(sourceCurrency, targetCurrency)
	if err != nil {
		return 0, fmt.Errorf("unable to get exchange rates: %w", err)
	}

	return exchangeRates, nil
}

func getECBExchangeRate(source, target currency) (changeRate, error) {
	const exchangeRatesURL = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	// build the HTTP request
	req, err := http.NewRequest(http.MethodGet, exchangeRatesURL, nil)
	if err != nil {
		return 0., fmt.Errorf("unable to build http request to %s: %w", exchangeRatesURL, err)
	}

	// use the default http client provided by the http library
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0., fmt.Errorf("error while requesting URL %s: %w", exchangeRatesURL, err)
	}

	// don't forget to close the response's body
	defer resp.Body.Close()

	if err = checkStatusCode(resp.StatusCode); err != nil {
		// TODO: have our own error?
		return 0., fmt.Errorf("invalid status code: %w", err)
	}

	// read the response
	decoder := xml.NewDecoder(resp.Body)
	var ecbMessage Envelope
	err = decoder.Decode(&ecbMessage)
	if err != nil {
		return 0., fmt.Errorf("unable to decode message: %w", err)
	}

	// do we want to returnt his directly ?
	rate, err := ecbMessage.changeRate(source, target)
	if err != nil {
		return 0., fmt.Errorf("couldn't find the exchange rate: %w", err)
	}

	return rate, nil
}
