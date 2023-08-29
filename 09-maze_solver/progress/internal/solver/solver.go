package solver

import (
	"fmt"
	"image"
	"image/png"
	"log/slog"
	"os"
)

// Solver is capable of finding the path through a maze.
type Solver struct {
	maze           *image.RGBA
	config         config
	pathsToExplore chan []point2d

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

// SaveSolution saves the image as a PNG file with the solution path in red.
func (s *Solver) SaveSolution(outputPath string) error {
	_, err := os.Stat(outputPath)
	switch {
	case err == nil:
		return fmt.Errorf("output file %s already exists", outputPath)
	case !os.IsNotExist(err):
		return fmt.Errorf("unable to check output file %s: %w", outputPath, err)
	}

	for _, p := range s.solution {
		s.maze.Set(p.x, p.y, s.config.solutionColour)
	}

	fd, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("unable to create output image file at %s", outputPath)
	}
	defer fd.Close()

	err = png.Encode(fd, s.maze)
	if err != nil {
		return fmt.Errorf("unable to write output image at %s", outputPath)
	}

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
