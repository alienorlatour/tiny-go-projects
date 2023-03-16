package books

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
	villette      = Book{Author: "Charlotte Brontë", Title: "Villette"}
	ilPrincipe    = Book{Author: "Niccolò Machiavelli", Title: "Il Principe"}
)

func noError(t *testing.T, err error) {
	t.Helper()
	assert.NoError(t, err)
}

func TestLoad(t *testing.T) {
	tests := map[string]struct {
		collsFile string
		want      Collectors
		checkErr  func(*testing.T, error)
	}{
		"file exists": {
			collsFile: "testdata/colls.json",
			want: Collectors{
				{Name: "Fadi", Items: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Items: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
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
			got, err := Load(testCase.collsFile)
			testCase.checkErr(t, err)
			assert.Equal(t, testCase.want, got)
		})
	}
}
