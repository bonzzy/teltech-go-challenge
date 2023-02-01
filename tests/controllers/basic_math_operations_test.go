package test

import (
	"encoding/json"
	"fmt"
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/core"
	"github.com/bonzzy/teltech-go-challenge/dtos"
	"github.com/bonzzy/teltech-go-challenge/tests/test_helpers"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type testCaseExpected struct {
	answer float64
	x      float64
	y      float64
	cached bool
}
type testCase struct {
	x                  string
	y                  string
	action             string
	expected           testCaseExpected
	expectedBadRequest bool
	message            string
}

func TestMathOperations(t *testing.T) {
	simpleTestCases := []testCase{
		{
			x:        "1.0",
			y:        "1.0",
			action:   "add",
			expected: testCaseExpected{answer: 2.0, x: 1.0, y: 1.0, cached: false},
			message:  "Simple add",
		},
		{
			x:        "2.0",
			y:        "1.0",
			action:   "subtract",
			expected: testCaseExpected{answer: 1.0, x: 2.0, y: 1.0, cached: false},
			message:  "Simple subtract",
		},
		{
			x:        "-2.3",
			y:        "-2.4",
			action:   "subtract",
			expected: testCaseExpected{answer: 0.1, x: -2.3, y: -2.4, cached: false},
			message:  "Float problem with subtract",
		},
		{
			x:        "4",
			y:        "-2",
			action:   "multiply",
			expected: testCaseExpected{answer: -8, x: 4, y: -2, cached: false},
			message:  "Simple multiply",
		},
	}

	edgeTestCases := []testCase{
		{
			x:        "100000000.0",
			y:        "1000000000.0",
			action:   "add",
			expected: testCaseExpected{answer: 1100000000.0, x: 100000000.0, y: 1000000000.0, cached: false},
			message:  "Large numbers",
		},
		{
			x:        "0.00001",
			y:        "0.000000001",
			action:   "add",
			expected: testCaseExpected{answer: 0.000010001, x: 0.00001, y: 0.000000001, cached: false},
			message:  "Small numbers",
		},
		{
			x:        "-10.0",
			y:        "5.0",
			action:   "subtract",
			expected: testCaseExpected{answer: -15, x: -10, y: 5, cached: false},
			message:  "Negative numbers",
		},
		{
			x:        "0.123456789",
			y:        "0.987654321",
			action:   "add",
			expected: testCaseExpected{answer: 1.111111110, x: 0.123456789, y: 0.987654321, cached: false},
			message:  "Floating-point precision",
		},
		{
			x:                  "10.0",
			y:                  "0.0",
			action:             "divide",
			expectedBadRequest: true,
			message:            "Division by zero",
		},
		{
			x:        "1000000000.0",
			y:        "0.1",
			action:   "multiply",
			expected: testCaseExpected{answer: 100000000.0, x: 1000000000.0, y: 0.1, cached: false},
			message:  "Overflow and underflow",
		},
		{
			x:        "0.12345678912345",
			y:        "0.987654321987654321",
			action:   "add",
			expected: testCaseExpected{answer: 1.1111111111111043, x: 0.12345678912345, y: 0.9876543219876543, cached: false},
			message:  "Decimal place rounding",
		},
	}
	testCases := append(simpleTestCases, edgeTestCases...)

	for _, testCase := range testCases {
		handle := getHandle(testCase.action)

		if handle == nil {
			panic(fmt.Sprintf("Handle for action %s not found!", testCase.action))
		}

		path := fmt.Sprintf("/%s?x=%s&y=%s", testCase.action, testCase.x, testCase.y)
		w := test_helpers.PerformRequest(handle, "GET", path)

		expected, err := json.Marshal(dtos.Response{Action: testCase.action, X: testCase.expected.x, Y: testCase.expected.y, Answer: testCase.expected.answer, Cached: testCase.expected.cached})

		if err != nil {
			panic(err)
		}

		if testCase.expectedBadRequest {
			assert.Equal(t, http.StatusBadRequest, w.Code)
			continue
		}

		// go json encoder terminates each value with a newline.
		// https://go.dev/src/encoding/json/stream.go?s=4272:4319
		assert.Equal(t, string(expected)+"\n", w.Body.String(), testCase.message)
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
