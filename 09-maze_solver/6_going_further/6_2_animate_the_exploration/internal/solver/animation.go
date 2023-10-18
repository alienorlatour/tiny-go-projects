package solver

import (
	"image"
	"image/color/palette"

	"golang.org/x/image/draw"
)

// countExplorablePixels scans the maze and counts the number
// of pixels that are not walls.
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

// registerExploredPixels registers positions as explored on the image,
// and, if we reach a threshold, adds the frame to the output GIF.
func (s *Solver) registerExploredPixels() {
	// totalExpectedFrames is the number of frames we want in the output gif.
	// We won't get exactly 30, because we won't be exploring every pixel.
	const totalExpectedFrames = 30

	explorablePixels := s.countExplorablePixels()
	pixelsExplored := 0

	for {
		select {
		case <-s.quit:
			return
		case pos := <-s.exploredPixels:
			s.maze.Set(pos.X, pos.Y, s.config.exploredColour)
			pixelsExplored++
			if pixelsExplored%(explorablePixels/totalExpectedFrames) == 0 {
				s.drawCurrentFrameToGIF()
			}
		}
	}
}

// drawCurrentFrameToGIF adds the current state of the maze as a frame of the animation.
func (s *Solver) drawCurrentFrameToGIF() {
	const (
		// gifSize is the length and width of the generated GIF.
		gifSize = 500
		// frameDuration is the duration in hundredth of a second of each frame.
		// 20 hundredths of a second per frame means 5 frames per second.
		frameDuration = 20
	)

	// Create a paletted frame that has the same ratio as the input image
	frame := image.NewPaletted(image.Rect(0, 0, gifSize, gifSize*s.maze.Bounds().Dy()/s.maze.Bounds().Dx()), palette.Plan9)

	// Convert RGBA to paletted
	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, frameDuration)
}