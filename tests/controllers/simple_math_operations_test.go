package test

import (
	"encoding/json"
	"fmt"
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	x := 1.0
	y := 1.0
	answer := 2.0
	path := fmt.Sprintf("/add?x=%f&y=%f", x, y)
	w := helpers.PerformRequest(controllers.Add, "GET", path)

	expected, err := json.Marshal(controllers.Response{Action: "add", X: x, Y: y, Answer: answer, Cached: false})

	if err != nil {
		panic(err)
	}

	assert.Equal(t, string(expected)+"\n", w.Body.String())
}
