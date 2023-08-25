package solver

import (
	"fmt"
	"log/slog"
	"slices"
)

const poolSize = 3

func (s *Solver) listenToBranches() {
	// create pool of N workers
	for i := 0; i < poolSize; i++ {
		go func() {
			for {
				path, ok := <-s.pathsToExplore
				if !ok {
					// the solution has been found
					return
				}

				// read a job to perform
				s.explore(path)
			}
		}()
	}
}

func (s *Solver) explore(pathToBranch []point2d) {
	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for {
		if s.solution != nil {
			return
		}

		s.explored <- pos

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
				s.pathsToExploreMutex.Lock()
				close(s.pathsToExplore)
				s.pathsToExploreMutex.Unlock()
				s.solution = append(pathToBranch, n)

				s.explored <- n
				close(s.explored)
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
				branch := append(slices.Clone(pathToBranch), candidates[i])
				slog.Info(fmt.Sprintf("%v seems to be promising", candidates[i]))
				s.pathsToExploreMutex.Lock()
				if s.solution != nil {
					return
				}
				s.pathsToExplore <- branch
				s.pathsToExploreMutex.Unlock()
			}
			// Continue exploration on this branch.
			pathToBranch = append(pathToBranch, candidates[0])
			previous = pos
			pos = candidates[0]
		default:
			slog.Error("whaaat?")
		}
	}
}
