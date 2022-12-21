package repository

import (
	"fmt"

	"github.com/ablqk/tiny-go-projects/chapter-05/layered/money"
)

type envelope struct {
	Cube envelopeCube `xml:"cube"`
}

type envelopeCube struct {
	ParentCube parentCube `xml:"cube"`
}

type parentCube struct {
	Time  string `xml:"time,attr"`
	Cubes []cube `xml:"cube"`
}

type cube struct {
	Currency string  `xml:"currency,attr"`
	Rate     float32 `xml:"rate,attr"`
}

// changeRate reads the change rate from the envelope's contents.
func (e envelope) changeRate(source, target money.Currency) (money.ChangeRate, error) {
	var foundSource, foundTarget bool
	var factor money.ChangeRate

	for _, c := range e.Cube.ParentCube.Cubes {
		switch c.Currency {
		case source.Code():
			factor = 1. / money.ChangeRate(c.Rate)
			foundSource = true
		case target.Code():
			factor = money.ChangeRate(c.Rate)
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
func (e envelope) Equal(other envelope) bool {
	return e.Cube.Equal(other.Cube)
}

// Equal tells whether the 2 EnvelopeCubes are equal.
func (ec envelopeCube) Equal(other envelopeCube) bool {
	return ec.ParentCube.Equal(other.ParentCube)
}

// Equal tells whether the 2 ParentCubes are equal.
func (pc parentCube) Equal(other parentCube) bool {
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
func (c cube) Equal(other cube) bool {
	if c.Currency != other.Currency {
		return false
	}
	// TODO: Add a tolerance?
	if c.Rate != other.Rate {
		return false
	}
	return true
}
