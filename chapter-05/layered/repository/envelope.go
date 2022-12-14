package repository

import (
	"fmt"

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

// changeRate reads the change rate from the Envelope's contents.
func (e Envelope) changeRate(source, target money.Currency) (money.ChangeRate, error) {
	var foundSource, foundTarget bool
	var factor money.ChangeRate

	for _, cube := range e.Cube.ParentCube.Cubes {
		switch cube.Currency {
		case source.Code():
			factor = 1. / money.ChangeRate(cube.Rate)
			foundSource = true
		case target.Code():
			factor = money.ChangeRate(cube.Rate)
			foundTarget = true
		}
	}

	// EUR is allowed not to be found
	if !foundSource && source.Code() != "EUR" {
		return 0., fmt.Errorf("unable to find source currency %s", source.Code())
	}

	if !foundTarget && target.Code() != "EUR" {
		return 0., fmt.Errorf("unable to find target currency %s", target.Code())
	}

	// TODO AL: what happens with CHF to CHF?
	// TODO AL: Why do we need the booleans?

	return factor, nil
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
