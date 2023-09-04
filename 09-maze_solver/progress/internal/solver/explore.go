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
	// A path starts at the entrance and has stepped into the maze
	// for at least 1 pixel, for a total of 2 pixels.
	if len(pathToBranch) < 2 {
		return
	}

	pos := pathToBranch[len(pathToBranch)-1]
	previous := pathToBranch[len(pathToBranch)-2]

	for s.solution == nil {
		s.markPixelExplored(pos)

		// We know we'll have between up to 3 new neighbours to explore.
		candidates := make([]image.Point, 0, 3)
		for _, n := range neighbours(pos) {
			if n == previous {
				continue
			}

			switch s.maze.RGBAAt(n.X, n.Y) {
			case s.config.treasureColour:
				slog.Info("Solution found!")
				s.mutex.Lock()
				close(s.quit)
				s.mutex.Unlock()
				s.solution = append(pathToBranch, n)
				return
			case s.config.pathColour:
				candidates = append(candidates, n)
			}
		}

		switch len(candidates) {
		case 0:
			// slog.Info(fmt.Sprintf("I must have taken the wrong turn at %v.", pos))
			return
		case 1, 2, 3:
			for i := 1; i < len(candidates); i++ {
				branch := append(slices.Clone(pathToBranch), candidates[i])
				s.mutex.Lock()
				// We are sure we send to pathsToExplore only when the quit channel isn't closed.
				select {
				// We're reading a zero-value when the channel is closed, otherwise we go to default.
				case <-s.quit:
					// Someone else has found the solution.
					slog.Info("I'm an unlucky branch, someone else found the treasure, I quit.")
					s.mutex.Unlock()
					return
				default:
					s.pathsToExplore <- branch
				}
				s.mutex.Unlock()
			}

			pathToBranch = append(pathToBranch, candidates[0])
			previous = pos
			pos = candidates[0]
		}
	}
}

// markPixelExplored registers a position as explored on the image,
// and, if we reach a threshold, adds the frame to the output GIF.
func (s *Solver) markPixelExplored(pos image.Point) {
	s.maze.Set(pos.X, pos.Y, s.config.exploredColour)
	currentCount := s.exploredCount.Add(1)
	if int(currentCount)%(s.maze.Bounds().Dx()*s.maze.Bounds().Dy())/200 == 0 {
		s.drawCurrentFrameToGif()
	}
}
