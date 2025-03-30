package getstatus

import (
	"gordle-service/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandle(t *testing.T) {
	// NOTE:
	// 1. We didn't supply any handler or start the server,
	// so how is the unit test working???
	req, err := http.NewRequest(http.MethodGet, "/games/", nil)
	require.NoError(t, err)

	req.SetPathValue(api.GameID, "123456")

	recorder := httptest.NewRecorder()

	Handle(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"sta tus":""}`, recorder.Body.String())
}
