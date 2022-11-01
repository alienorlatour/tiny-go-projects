package gordle

// ConfigFunc defines a configuration function for the Gordle.
type ConfigFunc func(g *Gordle) error

// WithMaxAttempts changes the maximum number of attempts (default is unlimited)
func WithMaxAttempts(maxAttempts int) ConfigFunc {
	return func(g *Gordle) error {
		g.maxAttempts = maxAttempts
		return nil
	}
}
