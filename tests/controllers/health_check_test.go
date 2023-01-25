package test

import (
	"encoding/json"
	"github.com/bonzzy/teltech-go-challenge/setup"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHealthCheck(t *testing.T) {
	router := setup.SetupRouter()

	w := performRequest(router, "GET", "/healthz")

	expected, err := json.Marshal(gin.H{"up": true})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
