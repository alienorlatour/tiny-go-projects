package solver

import (
	"fmt"
	"image"
	"log/slog"
	"slices"
	"sync"
)

func (s *Solver) listenToBranches() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for {
		select {
		case <-s.quit:
			slog.Info(fmt.Sprint("the solution has been found, worker going to sleep"))
			return
		case p := <-s.pathsToExplore:
			wg.Add(1)
			go func(p []image.Point) {
				defer wg.Done()

				s.explore(p)
			}(p)
		}
	}
}

// explore one path and publish to the s.pathsToExplore channel any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch []image.Point) {
	if len(pathToBranch) < 2 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for s.solution == nil {
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if n == previous {
				continue
			}

			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.config.treasureColour:
				slog.Info("Solution found!")
				s.quit <- struct{}{}
				s.solution = append(pathToBranch, n)
				return
			case s.config.pathColour:
				candidates = append(candidates, n)
			}
		}

		switch len(candidates) {
		case 0:
			slog.Info("I must have taken the wrong turn :(")
			return
		case 1, 2, 3:
			for i := 1; i < len(candidates); i++ {
				branch := append(slices.Clone(pathToBranch), candidates[i])
				s.pathsToExplore <- branch
			}

			pathToBranch = append(pathToBranch, candidates[0])
			previous = pos
			pos = candidates[0]
		}
	}
}
