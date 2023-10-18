package solver

import (
	"fmt"
	"image"
	"log/slog"
	"sync"
)

// listenToBranches creates a new routine for each branch published in s.pathsToExplore.
func (s *Solver) listenToBranches() {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for p := range s.pathsToExplore {
		wg.Add(1)
		go func(path *path) {
			defer wg.Done()
			s.explore(path)
		}(p)
		if s.solution != nil {
			return
		}
	}
}

// explore one path and publish to the s.pathsToExplore channel
// any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch *path) {
	if pathToBranch == nil {
		// This is a safety net. It should be used, but when it's needed, at least it's there.
		return
	}

	pos := pathToBranch.at

	for s.solution == nil {
		// We know we'll have up to 3 new neighbours to explore.
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if pathToBranch.isPreviousStep(n) {
				// Let's not return to the previous position
				continue
			}
			// Look at the colour of this pixel.
			// RGBAAt returns a color.RGBA{} zero value if the pixel is outside the bounds of the image.
			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.config.treasureColour:
				s.mutex.Lock()
				// Even though we're inside a loop, we are returning from the function here,
				// which makes it safe to defer the call to Unlock.
				defer s.mutex.Unlock()

				if s.solution == nil {
					s.solution = &path{previousStep: pathToBranch, at: n}
					slog.Info(fmt.Sprintf("Treasure found: %v!", s.solution.at))
				}
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
			branch := &path{previousStep: pathToBranch, at: candidate}
			s.pathsToExplore <- branch
		}

		pathToBranch = &path{previousStep: pathToBranch, at: candidates[0]}
		pos = candidates[0]
	}
}
