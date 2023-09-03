package solver

import (
	"fmt"
	"image"
	"log/slog"
	"sync"
)

// Solver is capable of finding the path through a maze.
type Solver struct {
	maze           *image.RGBA
	config         config
	pathsToExplore chan []image.Point
	quit           chan struct{}

	// mutex protecting the channels, ensuring we don't send a new path when we should quit
	mutex sync.Mutex

	solution []image.Point
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
		pathsToExplore: make(chan []image.Point),
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

	go func() { s.pathsToExplore <- []image.Point{entrance, {1, entrance.Y}} }()
	s.listenToBranches()

	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (image.Point, error) {
	height := s.maze.Bounds().Dy() - 1

	for y := 1; y <= height-1; y++ {
		if s.maze.RGBAAt(0, y) == s.config.entranceColour {
			return image.Point{0, y}, nil
		}
	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}
