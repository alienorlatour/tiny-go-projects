package solver

import (
	"fmt"
	"log/slog"
)

func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		slog.Info(fmt.Sprintf("exploring new path %v", p))
		go s.explore(p)
	}
}

func (s *Solver) explore(pathToBranch []point2d) {
	if len(pathToBranch) == 1 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for {
		// Peek in each direction for path pixels
		candidates := []point2d{}
		for _, n := range pos.neighbours() {
			if n == previous {
				continue
			}

			// Check each neighbour's type.
			// They can only be Wall, Path, Start, or End.
			switch s.maze.RGBAAt(n.x, n.y) {
			case s.config.EndColour:
				s.solution = append(pathToBranch, n)
				slog.Info("Solution found!")
				return
			case s.config.PathColour:
				candidates = append(candidates, n)
			}
		}

		switch len(candidates) {
		case 0:
			// This is a dead end.
			return
		case 1, 2, 3:
			// Notify there is a new branch to explore.
			// If there is only one branch, it's the continuation of the current path,
			// and we don't enter this loop in that case.
			for i := 1; i < len(candidates); i++ {
				branch := append(pathToBranch, candidates[i])
				s.pathsToExplore <- branch
			}

			// Continue exploration on this branch.
			pathToBranch = append(pathToBranch, candidates[0])
			previous = pos
			pos = candidates[0]
		}
	}
}
