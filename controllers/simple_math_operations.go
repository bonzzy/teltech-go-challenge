package controllers

import (
	"github.com/bonzzy/teltech-go-challenge/core"
	"math/bits"
	"net/http"
	"strconv"
)

// https://github.com/go-playground/validator
type MathParams struct {
	X float64 `json:"x" form:"x" binding:"required,number"`
	Y float64 `json:"y" form:"y" binding:"required,number"`
}

type Response struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}

// TODO implement cache, every cache hit needs to refresh the expiration time
// should expire anything that wasn't hit for a min
func Add(request core.Request) core.Response {
	var x, y float64
	if len(request.Query["x"]) > 0 {
		x, _ = strconv.ParseFloat(request.Query["x"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	if len(request.Query["y"]) > 0 {
		y, _ = strconv.ParseFloat(request.Query["y"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	params := MathParams{
		X: x,
		Y: y,
	}
	// Validation
	// return http.StatusBadRequest if validation fails

	response := Response{
		Action: "add",
		X:      params.X,
		Y:      params.Y,
		Answer: params.Y + params.X,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}

func Subtract(request core.Request) core.Response {
	var x, y float64
	if len(request.Query["x"]) > 0 {
		x, _ = strconv.ParseFloat(request.Query["x"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	if len(request.Query["y"]) > 0 {
		y, _ = strconv.ParseFloat(request.Query["y"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	params := MathParams{
		X: x,
		Y: y,
	}

	response := Response{
		Action: "subtract",
		X:      params.X,
		Y:      params.Y,
		Answer: params.X - params.Y,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}

func Multiply(request core.Request) core.Response {
	var x, y float64
	if len(request.Query["x"]) > 0 {
		x, _ = strconv.ParseFloat(request.Query["x"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	if len(request.Query["y"]) > 0 {
		y, _ = strconv.ParseFloat(request.Query["y"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	params := MathParams{
		X: x,
		Y: y,
	}

	response := Response{
		Action: "multiply",
		X:      params.X,
		Y:      params.Y,
		Answer: params.X * params.Y,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}

}

func Divide(request core.Request) core.Response {
	var x, y float64
	if len(request.Query["x"]) > 0 {
		x, _ = strconv.ParseFloat(request.Query["x"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	if len(request.Query["y"]) > 0 {
		y, _ = strconv.ParseFloat(request.Query["y"][0], bits.UintSize)
	} else {
		// throw
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	params := MathParams{
		X: x,
		Y: y,
	}

	response := Response{
		Action: "divide",
		X:      params.X,
		Y:      params.Y,
		Answer: params.X / params.Y,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}
