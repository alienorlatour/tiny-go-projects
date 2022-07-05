package gordle

import (
	"fmt"
	"io"
	"os"
)

func New(cfs ...ConfigFunc) (*Gordle, error) {
	g := &Gordle{
		reader:      os.Stdin,
		maxAttempts: -1,
	}
	for _, cf := range cfs {
		err := cf(g)
		if err != nil {
			return nil, fmt.Errorf("unable to apply config func: %w", err)
		}
	}
	return g, nil
}

func (g *Gordle) Play() {
	for {

	}
}

type Gordle struct {
	reader      io.Reader
	maxAttempts int
}
