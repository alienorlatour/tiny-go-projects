package solver

import (
	"fmt"
	"log/slog"
	"slices"
)

func (s *Solver) listenToBranches() {
	for p := range s.pathsToExplore {
		go s.explore(p)
	}
}

func (s *Solver) explore(pathToBranch pointsWithID) {
	if len(pathToBranch.points) == 1 {
		return
	}

	pos := pathToBranch.points[len(pathToBranch.points)-1]
	previous := pathToBranch.points[len(pathToBranch.points)-2]

	for {
		s.mutex.Lock()
		// mark pos as seen
		s.maze.Set(pos.x, pos.y, s.config.ExploredColour)
		s.mutex.Unlock()

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
				s.mutex.Lock()
				close(s.pathsToExplore)
				s.solution = append(pathToBranch.points, n)
				s.mutex.Unlock()
				slog.Info("Solution found!")
				return
			case s.config.PathColour:
				candidates = append(candidates, n)
			}
		}

		switch len(candidates) {
		case 0:
			// This is a dead end.
			slog.Info("I must have taken the wrong turn :(")
			return
		case 1, 2, 3:
			// Notify there is a new branch to explore.
			// If there is only one branch, it's the continuation of the current path,
			// and we don't enter this loop in that case.
			for i := 1; i < len(candidates); i++ {
				branch := append(slices.Clone(pathToBranch.points), candidates[i])
				slog.Info(fmt.Sprintf("%v (%s-%d) seems to be promising", candidates[i], pathToBranch.id, i))

				s.mutex.Lock()
				if s.solution != nil {
					return
				}
				s.pathsToExplore <- pointsWithID{branch, fmt.Sprintf("%s-%d", pathToBranch.id, i)}
				s.mutex.Unlock()
			}
			// Continue exploration on this branch.
			pathToBranch.points = append(pathToBranch.points, candidates[0])
			previous = pos
			pos = candidates[0]
		default:
			slog.Error("whaaat?")
		}
	}
}
