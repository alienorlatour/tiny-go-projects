package solver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOpenImage_errors(t *testing.T) {
	testCases := map[string]struct {
		input string
		err   string
	}{
		"no such file": {
			input: "nosuchfile.png",
			err:   "unable to check input file",
		},
		"not a png": {
			input: "file.go",
			err:   "unable to load input image",
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			img, err := openImage(tc.input)

			assert.Nil(t, img)
			assert.Error(t, err)
			assert.ErrorContains(t, err, tc.err)
		})
	}
}
