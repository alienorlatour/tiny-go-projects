package solver

import (
	"log/slog"
	"slices"
)

func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		go s.explore(p)
	}
}

// explore one path and publish to the s.pathsToExplore channel any branch we discover that we don't take.
func (s *Solver) explore(pathToBranch []point2d) {
	if len(pathToBranch) < 2 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for {
		candidates := make([]point2d, 0, 3)
		for _, n := range pos.neighbours() {
			if n == previous {
				continue
			}

			switch s.maze.RGBAAt(n.x, n.y) {
			case s.config.treasureColour:
				slog.Info("Solution found!")
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
