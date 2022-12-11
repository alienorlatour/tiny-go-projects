package money

import "fmt"

func (e Envelope) changeRate(source, target currency) (changeRate, error) {
	factor := changeRate(1.0)
	var foundSource, foundTarget bool
	for _, cube := range e.Cube.ParentCube.Cubes {
		switch cube.Currency {
		case source.code:
			// TODO: I don't really like this
			factor /= changeRate(cube.Rate)
			foundSource = true
		case target.code:
			// TODO: I don't really like this
			factor *= changeRate(cube.Rate)
			foundTarget = true
		}
	}
	// EUR is allowed not to be found
	if !foundSource && source.code != "EUR" {
		return 0., fmt.Errorf("unable to find source currency %s", source.code)
	}
	if !foundTarget && target.code != "EUR" {
		return 0., fmt.Errorf("unable to find target currency %s", target.code)
	}
	return factor, nil
}

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

func (e Envelope) Equal(other Envelope) bool {
	return e.Cube.Equal(other.Cube)
}

func (ee EnvelopeCube) Equal(other EnvelopeCube) bool {
	return ee.ParentCube.Equal(other.ParentCube)
}

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

func (c Cube) Equal(other Cube) bool {
	if c.Currency != other.Currency {
		return false
	}
	// TODO: Add a toleration?
	if c.Rate != other.Rate {
		return false
	}
	return true
}
