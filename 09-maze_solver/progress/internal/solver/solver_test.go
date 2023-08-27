package solver

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolver_findEntrance_success(t *testing.T) {
	tests := map[string]struct {
		inputPath string
		want      point2d
	}{
		"middle": {
			inputPath: "testdata/maze10_10.png",
			want:      point2d{0, 5},
		},
		"somewhere else": {
			inputPath: "testdata/maze15_15.png",
			want:      point2d{0, 7},
		},
	}
	for name, tt := range tests {
		name, tt := name, tt

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			img, err := openImage(tt.inputPath)
			require.NoError(t, err)

			s := &Solver{
				maze:   img,
				config: defaultColours(),
			}

			got, err := s.findEntrance()

			assert.Equalf(t, tt.want, got, "findEntrance()")
		})
	}
}
