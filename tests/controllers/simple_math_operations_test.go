package test

import (
	"encoding/json"
	"fmt"
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/tests/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCaseExpected struct {
	answer float64
	cached bool
}
type testCase struct {
	x        float64
	y        float64
	action   string
	expected testCaseExpected
}

func TestMathOperations(t *testing.T) {
	testCases := []testCase{
		{
			x:        1.0,
			y:        1.0,
			action:   "add",
			expected: testCaseExpected{answer: 2.0, cached: false},
		},
		{
			x:        1.0,
			y:        -500.0,
			action:   "add",
			expected: testCaseExpected{answer: -499.0, cached: false},
		},
	}

	for _, testCase := range testCases {
		path := fmt.Sprintf("/add?x=%f&y=%f", testCase.x, testCase.y)
		w := helpers.PerformRequest(controllers.Add, "GET", path)

		expected, err := json.Marshal(controllers.Response{Action: testCase.action, X: testCase.x, Y: testCase.y, Answer: testCase.expected.answer, Cached: testCase.expected.cached})

		if err != nil {
			panic(err)
		}

		assert.Equal(t, string(expected)+"\n", w.Body.String())
	}
}
