package gordle

import "io"

type ConfigFunc func(g *Gordle) error

func WithReader(reader io.Reader) ConfigFunc {
	return func(g *Gordle) error {
		g.reader = reader
		return nil
	}
}
