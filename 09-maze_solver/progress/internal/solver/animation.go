package solver

import (
	"image"
	"image/color/palette"

	"golang.org/x/image/draw"
)

const (
	// totalExpectedFrames is the number of frames we want in the output gif.
	// We won't get exactly 30, because we won't be exploring every pixel.
	totalExpectedFrames = 30
	// gifSize is the length and width of the generated GIF.
	gifSize = 500
)

func (s *Solver) drawFrames() {
	pixelsExplored := 0
	explorablePixels := s.countExplorablePixels()

	for {
		select {
		case pos := <-s.exploredPixels:
			s.maze.Set(pos.X, pos.Y, s.config.exploredColour)
			pixelsExplored++
			if pixelsExplored%(explorablePixels/totalExpectedFrames) == 0 {
				s.drawCurrentFrameToGIF()
			}
		case <-s.quit:
			return
		}
	}
}

// countExplorablePixels scans the maze and counts the number of pixels that are not walls.
func (s *Solver) countExplorablePixels() int {
	explorablePixels := 0
	for row := 0; row < s.maze.Bounds().Dy(); row++ {
		for col := 0; col < s.maze.Bounds().Dx(); col++ {
			if s.maze.RGBAAt(col, row) != s.config.wallColour {
				explorablePixels++
			}
		}
	}
	return explorablePixels
}

// drawCurrentFrameToGIF adds the current state of the maze as a frame of the animation.
func (s *Solver) drawCurrentFrameToGIF() {
	// Create a paletted frame that has the same ratio as the input image
	frame := image.NewPaletted(image.Rect(0, 0, gifSize, gifSize*s.maze.Bounds().Dy()/s.maze.Bounds().Dx()), palette.Plan9)

	// Convert RGBA to paletted
	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, 1 /* hundredth of a second */)
}
