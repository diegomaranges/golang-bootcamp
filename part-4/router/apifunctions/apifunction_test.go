package apifunctions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func test(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://example.com/cars/1/13", nil)
	w := httptest.NewRecorder()
	ReturnItem(w, req)
	resp := w.Result()

	assert.Equal(t, resp.StatusCode, http.StatusNotFound, "found element")
}
