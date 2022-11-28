package money

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
	if c.Rate != other.Rate {
		return false
	}
	if c.Currency != other.Currency {
		return false
	}
	return true
}
