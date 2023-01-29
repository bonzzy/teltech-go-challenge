package test

import (
	"encoding/json"
	"fmt"
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/core"
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
	simpleTestCases := []testCase{
		{
			x:        1.0,
			y:        1.0,
			action:   "add",
			expected: testCaseExpected{answer: 2.0, cached: false},
		},
		{
			x:        1.0,
			y:        -500.5,
			action:   "add",
			expected: testCaseExpected{answer: -499.5, cached: false},
		},
		{
			x:        2.0,
			y:        1.0,
			action:   "subtract",
			expected: testCaseExpected{answer: 1.0, cached: false},
		},
		// test float problem
		{
			x:        -2.3,
			y:        -2.4,
			action:   "subtract",
			expected: testCaseExpected{answer: 0.1, cached: false},
		},
	}
	testCases := append(simpleTestCases)

	for _, testCase := range testCases {
		handle := getHandle(testCase.action)

		if handle == nil {
			panic(fmt.Sprintf("Handle for action %s not found!", testCase.action))
		}

		path := fmt.Sprintf("/%s?x=%f&y=%f", testCase.action, testCase.x, testCase.y)
		w := helpers.PerformRequest(handle, "GET", path)

		expected, err := json.Marshal(controllers.Response{Action: testCase.action, X: testCase.x, Y: testCase.y, Answer: testCase.expected.answer, Cached: testCase.expected.cached})

		if err != nil {
			panic(err)
		}

		assert.Equal(t, string(expected)+"\n", w.Body.String())
	}
}

func getHandle(action string) core.Handler[any] {
	switch action {
	case "add":
		return controllers.Add
	case "subtract":
		return controllers.Subtract
	case "multiply":
		return controllers.Multiply
	case "divide":
		return controllers.Divide
	}

	return nil
}
