package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"learngo-pockets/templates/internal/handlers/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer_assets(t *testing.T) {
	rr := httptest.NewRecorder()

	cli := mocks.NewHabitsClientMock(t)

	s := New(cli, t)
	s.assets(rr, httptest.NewRequest(http.MethodGet, "/assets/styles.css", nil))

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	expect, err := os.ReadFile("testdata/index_emptylist.html")
	require.NoError(t, err)
	assert.Equal(t, string(expect), rr.Body.String())
}
