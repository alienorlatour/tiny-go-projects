package solver

import (
	"fmt"
	"image"
	"log/slog"
	"sync"
)

// Solver is capable of finding the path from the entrance to the treasure.
// The maze has to be a RGBA image.
type Solver struct {
	maze           *image.RGBA
	config         config
	pathsToExplore chan *path

	solution *path
	mutex    sync.Mutex
}

// New builds a Solver by taking the path to the PNG maze, encoded in RGBA.
func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze:           img,
		config:         defaultColours(),
		pathsToExplore: make(chan *path, 1),
	}, nil
}

// Solve finds the path from the entrance to the treasure.
func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("unable to find entrance: %w", err)
	}

	slog.Info(fmt.Sprintf("starting at %v", entrance))

	// Write once in pathsToExplore before starting listening on the channel.
	s.pathsToExplore <- &path{previousSteps: nil, at: entrance}
	s.listenToBranches()

	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (image.Point, error) {
	for row := 0; row < s.maze.Bounds().Dy(); row++ {
		for col := 0; col < s.maze.Bounds().Dy(); col++ {
			if s.maze.RGBAAt(col, row) == s.config.entranceColour {
				return image.Point{X: col, Y: row}, nil
			}
		}
	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}
