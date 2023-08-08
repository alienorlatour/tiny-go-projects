package solver

import "sync"

type workerPool struct {
	sem chan struct{}

	p sync.Pool
}
