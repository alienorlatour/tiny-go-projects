package solver

import (
	"fmt"
	"image"
	"log/slog"
	"slices"
)

// listenToBranches creates new routine for each new branch published in s.pathsToExplore.
func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		go s.explore(p)
		if len(s.solution) != 0 {
			return
		}
	}

}

// explore one path and publish to the s.pathsToExplore channel
// any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch []image.Point) {
	// A path starts at the entrance and has stepped into the maze
	// for at least 1 pixel, for a total of 2 pixels minimum.
	if len(pathToBranch) < 2 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for s.solution == nil {
		// We know we'll have between up to 3 new neighbours to explore.
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if n == previous {
				continue
			}

			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.config.treasureColour:
				s.solution = append(pathToBranch, n)
				slog.Info(fmt.Sprintf("Treasure found: %v!", s.solution))
				return
			case s.config.pathColour:
				candidates = append(candidates, n)
			}
		}

		switch len(candidates) {
		case 0:
			slog.Info(fmt.Sprintf("I must have taken the wrong turn at %v.", pos))
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
