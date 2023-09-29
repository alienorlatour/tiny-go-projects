package solver

import (
	"image"
	"image/color/palette"

	"golang.org/x/image/draw"
)

const (
	// pathRatio is the proportion of path pixels in the total image, which consists of path + wall + treasure + entrance.
	// This value is very approximate, it would be between 0.25 and 0.5, depending on the maze.
	pathRatio = 0.4
	// totalExpectedFrames is the number of frames we want in the output gif.
	// We won't get exactly 30, because pathRatio is approximate. But we'll get something around 30.
	totalExpectedFrames = 30
	// gifSize is the length and width of the generated GIF.
	gifSize = 500
)

func (s *Solver) drawFrames() {
	pixelsExplored := 0
	totalPixels := s.maze.Bounds().Dx() * s.maze.Bounds().Dy()
	explorablePixels := int(float32(totalPixels) * pathRatio)

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

// drawCurrentFrameToGIF adds the current state of the maze as a frame of the animation.
func (s *Solver) drawCurrentFrameToGIF() {
	frame := image.NewPaletted(image.Rect(0, 0, gifSize, gifSize), palette.Plan9)

	// Convert RGBA to paletted
	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, 1 /* hundredth of a second */)
}
