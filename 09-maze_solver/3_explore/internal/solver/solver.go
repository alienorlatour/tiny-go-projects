package solver

import (
	"fmt"
	"image"
	"log/slog"
)

// Solver is capable of finding the path from the entrance to the treasure.
// The maze has to be a RGBA image.
type Solver struct {
	maze           *image.RGBA
	config         config
	pathsToExplore chan []image.Point

	solution []image.Point
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
		pathsToExplore: make(chan []image.Point, 1),
	}, nil
}

// Solve finds the path from the entrance to the treasure.
func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("unable to find entrance: %w", err)
	}

	slog.Info(fmt.Sprintf("starting at %v", entrance))

	// The first pixel is on the edge, the second pixel is inwards.
	s.pathsToExplore <- []image.Point{entrance, {1, entrance.Y}}

	// Listen for new paths to explore. This only returns when the maze is solved.
	s.listenToBranches()

	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (image.Point, error) {
	height := s.maze.Bounds().Dy()

	for y := 1; y < height-1; y++ {
		if s.maze.RGBAAt(0, y) == s.config.entranceColour {
			return image.Point{0, y}, nil
		}

	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}