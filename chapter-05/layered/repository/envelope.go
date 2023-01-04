package repository

import (
	"errors"
	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

type Envelope struct {
	Cube EnvelopeCube `xml:"Cube"`
}

type EnvelopeCube struct {
	ParentCube ParentCube `xml:"Cube"`
}

type ParentCube struct {
	Time  string `xml:"time,attr"`
	Cubes []Cube `xml:"Cube"`
}

type Cube struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

func (e Envelope) loadChangeRates() map[string]float32 {
	changeRates := make(map[string]float32)
	for _, c := range e.Cube.ParentCube.Cubes {
		changeRates[c.Currency] = c.Rate
	}

	// default ecb has EUR to x currency
	changeRates["EUR"] = 1.

	return changeRates
}

// changeRate reads the change rate from the Envelope's contents.
func (e Envelope) changeRate(source, target money.Currency) (money.ChangeRate, error) {
	if source == target {
		// No change rate for same source and target currencies.
		return 1., nil
	}

	// changeRates stores the rates when Envelope is parsed.
	changeRates := e.loadChangeRates()

	sourceFactor, sourceFound := changeRates[source.Code()]
	targetFactor, targetFound := changeRates[target.Code()]

	if !sourceFound {
		return 0, errors.New("failed to found the source currency")
	}

	if !targetFound {
		return 0, errors.New("failed to found target currency")

	}

	return money.ChangeRate(targetFactor / sourceFactor), nil
}

// Equal tells whether the 2 Envelopes are equal.
func (e Envelope) Equal(other Envelope) bool {
	return e.Cube.Equal(other.Cube)
}

// Equal tells whether the 2 EnvelopeCubes are equal.
func (ec EnvelopeCube) Equal(other EnvelopeCube) bool {
	return ec.ParentCube.Equal(other.ParentCube)
}

// Equal tells whether the 2 ParentCubes are equal.
func (pc ParentCube) Equal(other ParentCube) bool {
	if pc.Time != other.Time {
		return false
	}
	if len(pc.Cubes) != len(other.Cubes) {
		return false
	}
	for i := range pc.Cubes {
		if !pc.Cubes[i].Equal(other.Cubes[i]) {
			return false
		}
	}
	return true
}

// Equal tells whether the 2 Cubes are equal.
func (c Cube) Equal(other Cube) bool {
	if c.Currency != other.Currency {
		return false
	}
	// TODO: Add a tolerance?
	if c.Rate != other.Rate {
		return false
	}
	return true
}
