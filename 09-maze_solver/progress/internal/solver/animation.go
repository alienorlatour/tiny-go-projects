package solver

import (
	"image"
	"image/color/palette"

	"golang.org/x/image/draw"
)

func (s *Solver) drawCurrentFrameToGif() {
	frame := image.NewPaletted(image.Rect(0, 0, 500, 500), palette.Plan9)

	// Convert RGBA to paletted
	draw.NearestNeighbor.Scale(frame, frame.Rect, s.maze, s.maze.Bounds(), draw.Over, nil)

	s.animation.Image = append(s.animation.Image, frame)
	s.animation.Delay = append(s.animation.Delay, 1)
}
