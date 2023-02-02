package ecbank

import (
	"errors"

	"github.com/ablqk/tiny-go-projects/chapter-05/final/money"
)

const baseCurrencyCode = "EUR"

type Envelope struct {
	Rates []CurrencyRate `xml:"Cube>Cube>Cube"`
}

type CurrencyRate struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

func (e Envelope) loadChangeRates() map[string]float32 {
	changeRates := make(map[string]float32)
	for _, c := range e.Rates {
		changeRates[c.Currency] = c.Rate
	}

	// default ecb has EUR to x currency
	changeRates[baseCurrencyCode] = 1.

	return changeRates
}

// changeRate reads the change rate from the Envelope's contents.
func (e Envelope) changeRate(source, target money.Currency) (money.ExchangeRate, error) {
	if source == target {
		// No change rate for same source and target currencies.
		return 1., nil
	}

	// changeRates stores the rates when Envelope is parsed.
	changeRates := e.loadChangeRates()

	sourceFactor, sourceFound := changeRates[source.Code()]
	if !sourceFound {
		return 0, errors.New("failed to found the source currency")
	}

	targetFactor, targetFound := changeRates[target.Code()]
	if !targetFound {
		return 0, errors.New("failed to found target currency")
	}

	return money.ExchangeRate(targetFactor / sourceFactor), nil
}
