package solver

import (
	"fmt"
	"image"
	"image/gif"
	"log"
	"sync"
)

// Solver is capable of finding the path from the entrance to the treasure.
// The maze has to be a RGBA image.
type Solver struct {
	maze    *image.RGBA
	palette palette

	pathsToExplore chan *path
	quit           chan struct{}

	solution *path
	mutex    sync.Mutex

	exploredPixels chan image.Point
	animation      *gif.GIF
}

// New builds a Solver by taking the path to the PNG maze, encoded in RGBA.
func New(imagePath string) (*Solver, error) {
	img, err := openMaze(imagePath)
	if err != nil {
		return nil, fmt.Errorf("cannot open maze image: %w", err)
	}

	return &Solver{
		maze:           img,
		palette:        defaultPalette(),
		pathsToExplore: make(chan *path, 1),
		quit:           make(chan struct{}),
		exploredPixels: make(chan image.Point),
		animation:      &gif.GIF{},
	}, nil
}

// Solve finds the path from the entrance to the treasure.
func (s *Solver) Solve() error {
	entrance, err := s.findEntrance()
	if err != nil {
		return fmt.Errorf("unable to find entrance: %w", err)
	}

	log.Printf("starting at %v", entrance)

	s.pathsToExplore <- &path{previousStep: nil, at: entrance}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Launch the goroutine in charge of drawing the GIF image.
		s.registerExploredPixels()
	}()

	go func() {
		defer wg.Done()
		// Listen for new paths to explore. This only returns when the maze is solved.
		s.listenToBranches()
	}()

	// Synchronisation step waiting for both goroutines to end
	wg.Wait()

	s.writeLastFrame()

	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (image.Point, error) {
	for row := s.maze.Bounds().Min.Y; row < s.maze.Bounds().Max.Y; row++ {
		for col := s.maze.Bounds().Min.X; col < s.maze.Bounds().Max.X; col++ {
			if s.maze.RGBAAt(col, row) == s.palette.entrance {
				return image.Point{X: col, Y: row}, nil
			}
		}
	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}
