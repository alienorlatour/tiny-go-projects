package ecbank

import (
	"encoding/xml"
	"fmt"
	"io"

	"learngo-pockets/moneyconverter/money"
)

const baseCurrencyCode = "EUR"

func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {
	// read the response
	decoder := xml.NewDecoder(respBody)

	var ecbMessage envelope
	err := decoder.Decode(&ecbMessage)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}

	rate, err := ecbMessage.changeRate(source, target)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}
	return rate, nil
}

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string             `xml:"currency,attr"`
	Rate     money.ExchangeRate `xml:"rate,attr"`
}

// mappedChangeRates builds a map of all the supported exchange rates.
func (e envelope) mappedChangeRates() map[string]money.ExchangeRate {
	changeRates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	for _, c := range e.Rates {
		changeRates[c.Currency] = c.Rate
	}

	// add EUR to EUR rate
	changeRates[baseCurrencyCode] = 1.

	return changeRates
}

// changeRate reads the change rate from the Envelope's contents.
func (e envelope) changeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		// No change rate for same source and target currencies.
		return 1., nil
	}

	// changeRates stores the rates when Envelope is parsed.
	changeRates := e.mappedChangeRates()

	sourceFactor, sourceFound := changeRates[source]
	if !sourceFound {
		return 0, fmt.Errorf("failed to find the source currency %s", source)
	}

	targetFactor, targetFound := changeRates[target]
	if !targetFound {
		return 0, fmt.Errorf("failed to find target currency %s", target)
	}

	return targetFactor / sourceFactor, nil
}
