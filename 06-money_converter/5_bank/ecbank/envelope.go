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

	var xrefMessage envelope
	err := decoder.Decode(&xrefMessage)
	if err != nil {
		return 0., fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}

	rate, err := xrefMessage.exchangeRate(source, target)
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

// exchangeRates builds a map of all the supported exchange rates.
func (e envelope) exchangeRates() map[string]money.ExchangeRate {
	rates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}

	// add EUR to EUR rate
	rates[baseCurrencyCode] = 1.

	return rates
}

// exchangeRate reads the change rate from the Envelope's contents.
func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		// No change rate for same source and target currencies.
		return 1., nil
	}

	// rates stores the rates when Envelope is parsed.
	rates := e.exchangeRates()

	sourceFactor, sourceFound := rates[source]
	if !sourceFound {
		return 0, fmt.Errorf("failed to find the source currency %s", source)
	}

	targetFactor, targetFound := rates[target]
	if !targetFound {
		return 0, fmt.Errorf("failed to find target currency %s", target)
	}

	return targetFactor / sourceFactor, nil
}
