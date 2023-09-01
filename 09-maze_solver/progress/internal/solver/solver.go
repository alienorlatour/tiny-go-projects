package solver

import (
	"fmt"
	"image"
	"log/slog"
)

// Solver is capable of finding the path through a maze.
type Solver struct {
	maze           *image.RGBA
	config         config
	pathsToExplore chan []point2d
	quit           chan struct{}

	solution []point2d
}

// New builds a Solver by taking the path to the PNG maze.
func New(imagePath string) (*Solver, error) {
	img, err := openImage(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze:           img,
		config:         defaultColours(),
		pathsToExplore: make(chan []point2d, 1),
		quit:           make(chan struct{}, 1),
	}, nil
}

// Solve finds the path from one end to the other.
func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("unable to find entrance: %w", err)
	}

	slog.Info(fmt.Sprintf("starting at %v", entrance))

	s.pathsToExplore <- []point2d{entrance, {1, entrance.y}}
	s.listenToBranches()

	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (point2d, error) {
	height := s.maze.Bounds().Dy() - 1

	for y := 1; y <= height-1; y++ {
		if s.maze.RGBAAt(0, y) == s.config.entranceColour {
			return point2d{0, y}, nil
		}
	}

	return point2d{}, fmt.Errorf("entrance position not found")
}
