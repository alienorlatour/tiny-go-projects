package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_index(t *testing.T) {
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	s := New(t)
	s.index(rr, nil)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	expect, err := os.ReadFile("testdata/index.html")
	require.NoError(t, err)
	assert.Equal(t, string(expect), rr.Body.String())
}
