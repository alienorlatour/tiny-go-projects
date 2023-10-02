package solver

import (
	"fmt"
	"image"
	"image/gif"
	"log/slog"
	"sync"
)

// Solver is capable of finding the path from the entrance to the treasure.
// The maze has to be a RGBA image.
type Solver struct {
	maze   *image.RGBA
	config config

	pathsToExplore chan *Path
	quit           chan struct{}

	exploredPixels chan image.Point
	animation      *gif.GIF

	solution *Path
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
		pathsToExplore: make(chan *Path),
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

	slog.Info(fmt.Sprintf("starting at %v", entrance))

	go func() {
		// The first pixel is on the edge, the second pixel is inwards.
		s.pathsToExplore <- &Path{PreviousSteps: nil, At: entrance}
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		// Launch the goroutine in charge of drawing the GIF image.
		s.drawFrames()
	}()

	go func() {
		defer wg.Done()
		// Listen for new paths to explore. This only returns when the maze is solved.
		s.listenToBranches()
	}()

	wg.Wait()

	s.finalise()

	return nil
}

// findEntrance returns the position of the maze entrance on the image.
func (s *Solver) findEntrance() (image.Point, error) {
	height := s.maze.Bounds().Dy()

	// We built the maze with the entrance on the left border, and not in a corner
	for y := 1; y < height-1; y++ {
		if s.maze.RGBAAt(0, y) == s.config.entranceColour {
			return image.Point{0, y}, nil
		}
	}

	return image.Point{}, fmt.Errorf("entrance position not found")
}

func (s *Solver) finalise() {
	stepsFromTreasure := s.solution
	// Paint the path from entrance to the treasure.
	for stepsFromTreasure != nil {
		s.maze.Set(stepsFromTreasure.At.X, stepsFromTreasure.At.Y, s.config.solutionColour)
		stepsFromTreasure = stepsFromTreasure.PreviousSteps
	}

	// Add the solution frame, with the coloured path, to the output gif.
	s.drawCurrentFrameToGIF()
	// Have the final frame containing the solution displayed for 3 seconds
	s.animation.Delay[len(s.animation.Delay)-1] = 300 /* hundredth of a second */
}
