package solver

import (
	"fmt"
	"log/slog"
	"slices"
	"sync"
)

const poolSize = 3

func (s *Solver) listenToBranches() {
	wg := sync.WaitGroup{}
	// create pool of N workers
	for i := 0; i < poolSize; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			slog.Info(fmt.Sprintf("starting worker %d", workerID))
			for {
				select {
				case <-s.quit:
					slog.Info(fmt.Sprintf("the solution has been found, worker %d going to sleep", workerID))
					return
				case path := <-s.pathsToExplore:
					slog.Info(fmt.Sprintf("[%d] found a new pos (%d in chan): %v", workerID, len(s.pathsToExplore), path))
					// read a job to perform
					s.explore(path)
				}
				//	slog.Info(fmt.Sprintf("worker %d has finished its job and is eligible for new work", workerID))
			}
		}(i)
	}

	wg.Wait()
}

func (s *Solver) explore(pathToBranch []point2d) {
	if len(pathToBranch) < 2 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for {
		s.toPaint <- pos

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
				s.b.broadcast(struct{}{})
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
				if s.solution != nil {
					return
				}
				slog.Info(fmt.Sprintf("queuing candidate"))
				s.pathsToExplore <- branch
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
