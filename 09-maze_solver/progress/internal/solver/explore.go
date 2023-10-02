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
		// s.quit will never return a value, unless something writes in it (which we don't do)
		// or it has been closed, which we do when we find the treasure.
		case <-s.quit:
			slog.Info(fmt.Sprint("the treasure has been found, worker going to sleep"))
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
	// A path starts at the entrance and has stepped into the maze
	// for at least 1 pixel, for a total of 2 pixels minimum.
	if len(pathToBranch) < 2 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for {
		// Let's first check whether we should quit.
		select {
		case <-s.quit:
			return
		case s.exploredPixels <- pos:
		}

		// We know we'll have up to 3 new neighbours to explore.
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if n == previous {
				continue
			}

			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.config.treasureColour:
				s.solution = append(pathToBranch, n)

				// Paint the path from entrance to the treasure.
				for _, p := range s.solution {
					// TODO: Data race here. The final pixel is sent to explore (l.48) and we paint it here (before closing).
					s.maze.Set(p.X, p.Y, s.config.solutionColour)
				}

				slog.Info(fmt.Sprintf("Treasure found: %v!", s.solution))
				close(s.quit)
				return
			case s.config.pathColour:
				candidates = append(candidates, n)
			}
		}

		if len(candidates) == 0 {
			slog.Debug("I must have taken the wrong turn.", "position", pos)
			return
		}

		for _, candidate := range candidates[1:] {
			branch := append(slices.Clone(pathToBranch), candidate)
			// We are sure we send to pathsToExplore only when the quit channel isn't closed.
			// A goroutine might have found the treasure since the check at the start of the loop.
			select {
			// s.quit returns a zero value only when the channel was closed, here -- when the exploration should end.
			case <-s.quit:
				slog.Debug("I'm an unlucky branch, someone else found the treasure, I give up.", "position", pos)
				return
			case s.pathsToExplore <- branch:
			}
		}

		pathToBranch = append(pathToBranch, candidates[0])
		previous = pos
		pos = candidates[0]
	}
}
