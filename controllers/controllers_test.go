package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gjlanc65/gintest/routes"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
