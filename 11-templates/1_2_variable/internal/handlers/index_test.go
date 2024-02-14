package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
)

func TestServer_index(t *testing.T) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	s := Server{}
	s.index(rr, nil)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	want, err := os.ReadFile("testdata/index.html")
	require.NoError(t, err)
	compareHTML(t, string(want), rr.Body.String())
}

func compareHTML(t *testing.T, left string, right string) {
	t.Helper()

	hleft, err := html.Parse(strings.NewReader(left))
	assert.NoError(t, err)
	hright, err := html.Parse(strings.NewReader(right))
	assert.NoError(t, err)

	assert.Equal(t, hleft.Data, hright.Data)
}
