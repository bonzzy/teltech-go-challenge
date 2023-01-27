package test

import (
	"encoding/json"
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	w := helpers.PerformRequest(controllers.Healthz, "GET", "/healthz")

	expected, err := json.Marshal(controllers.AppHealth{Up: true})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, 200, w.Code)
	// go json encoder terminates each value with a newline.
	// https://go.dev/src/encoding/json/stream.go?s=4272:4319
	assert.Equal(t, string(expected)+"\n", w.Body.String())
}
