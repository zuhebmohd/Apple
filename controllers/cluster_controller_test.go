package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetClusters(t *testing.T) {
	router := gin.Default()
	router.GET("/clusters", GetClusters)

	w := performRequest(router, "GET", "/clusters")
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Amazon")
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
