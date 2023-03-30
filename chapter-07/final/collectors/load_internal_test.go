package collectors

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func noError(t *testing.T, err error) {
	t.Helper()
	assert.NoError(t, err)
}

func TestLoad(t *testing.T) {
	tests := map[string]struct {
		collsFile string
		want      Collectors[string]
		checkErr  func(*testing.T, error)
	}{
		"file exists": {
			collsFile: "testdata/genericitems.json",
			want: Collectors[string]{
				{Name: "Fadi", Items: []string{"The Handmaid's Tale", "The Bell Jar"}},
				{Name: "Peggy", Items: []string{"Oryx and Crake", "The Handmaid's Tale", "Jane Eyre"}},
			},
			checkErr: noError,
		},
		"file doesn't exist": {
			collsFile: "testdata/no_file_here.json",
			want:      nil,
			checkErr: func(t *testing.T, err error) {
				pathErr := &os.PathError{}
				assert.ErrorAs(t, err, &pathErr)
				assert.Equal(t, "testdata/no_file_here.json", pathErr.Path)
			},
		},
		"invalid JSON": {
			collsFile: "testdata/invalid.json",
			want:      nil,
			checkErr: func(t *testing.T, err error) {
				jsonErr := &json.SyntaxError{}
				assert.ErrorAs(t, err, &jsonErr)
				assert.Equal(t, int64(174), jsonErr.Offset)
			},
		},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := Load[string](testCase.collsFile)
			testCase.checkErr(t, err)
			assert.Equal(t, testCase.want, got)
		})
	}
}
