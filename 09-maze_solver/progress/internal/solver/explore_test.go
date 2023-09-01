package solver

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolver_explore(t *testing.T) {
	tests := map[string]struct {
		inputImage string
		wantSize   int
	}{
		"cross": {
			inputImage: "testdata/explore_cross.png",
			wantSize:   2,
		},
		"dead end": {
			inputImage: "testdata/explore_deadend.png",
			wantSize:   0,
		},
		"double": {
			inputImage: "testdata/explore_double.png",
			wantSize:   1,
		},
		"treasure": {
			inputImage: "testdata/explore_treasure.png",
			wantSize:   0,
		},
		"treasure only": {
			inputImage: "testdata/explore_treasureonly.png",
			wantSize:   0,
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			maze, err := openImage(tt.inputImage)
			require.NoError(t, err)

			s := &Solver{
				maze:           maze,
				config:         defaultColours(),
				pathsToExplore: make(chan []image.Point, 3),
			}

			s.explore([]image.Point{{0, 2}, {1, 2}})

			assert.Equal(t, tt.wantSize, len(s.pathsToExplore))
		})
	}
}
