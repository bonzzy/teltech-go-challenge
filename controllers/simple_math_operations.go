package controllers

import (
	"github.com/bonzzy/teltech-go-challenge/core"
	"math/big"
	"net/http"
)

type MathParams struct {
	X *big.Float `json:"x" form:"x" binding:"required,number"`
	Y *big.Float `json:"y" form:"y" binding:"required,number"`
}

type ResponseBigFloat struct {
	Action string     `json:"action"`
	X      *big.Float `json:"x"`
	Y      *big.Float `json:"y"`
	Answer float64    `json:"answer"`
	Cached bool       `json:"cached"`
}

type Response struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}

func ValidateQueryXY(query map[string][]string) (bool, string) {
	if len(query["x"]) == 0 || !core.IsNumber(query["x"][0]) {
		return false, "x needs to be a number!"
	}
	if len(query["y"]) == 0 || !core.IsNumber(query["y"][0]) {
		return false, "y needs to be a number!"
	}

	return true, ""
}

func GetBigXYFromQuery(query map[string][]string) (*big.Float, *big.Float) {
	x, _, _ := new(big.Float).SetPrec(256).Parse(query["x"][0], 10)
	y, _, _ := new(big.Float).SetPrec(256).Parse(query["y"][0], 10)

	return x, y
}

// TODO implement cache, every cache hit needs to refresh the expiration time
// should expire anything that wasn't hit for a min
func Add(request core.Request) core.Response {
	roundDecimalPrecision := 16.0

	validation, err := ValidateQueryXY(request.Query)

	if !validation {
		return core.Response{HttpStatus: http.StatusBadRequest, Data: err}
	}

	x, y := GetBigXYFromQuery(request.Query)

	answer, _ := new(big.Float).SetPrec(256).Add(x, y).SetPrec(256).Float64()
	givenX, _ := new(big.Float).Set(x).SetPrec(256).Float64()
	givenY, _ := new(big.Float).Set(y).SetPrec(256).Float64()

	response := Response{
		Action: "add",
		X:      core.Round(givenX, roundDecimalPrecision),
		Y:      core.Round(givenY, roundDecimalPrecision),
		Answer: core.Round(answer, roundDecimalPrecision),
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}

func Subtract(request core.Request) core.Response {
	validation, err := ValidateQueryXY(request.Query)

	if !validation {
		return core.Response{HttpStatus: http.StatusBadRequest, Data: err}
	}

	x, y := GetBigXYFromQuery(request.Query)

	answer, _ := new(big.Float).SetPrec(256).Sub(x, y).SetPrec(256).Float64()
	givenX, _ := new(big.Float).Set(x).SetPrec(256).Float64()
	givenY, _ := new(big.Float).Set(y).SetPrec(256).Float64()

	response := Response{
		Action: "subtract",
		X:      givenX,
		Y:      givenY,
		Answer: answer,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}

func Multiply(request core.Request) core.Response {
	validation, err := ValidateQueryXY(request.Query)

	if !validation {
		return core.Response{HttpStatus: http.StatusBadRequest, Data: err}
	}

	x, y := GetBigXYFromQuery(request.Query)

	answer, _ := new(big.Float).SetPrec(256).Mul(x, y).SetPrec(256).Float64()
	givenX, _ := new(big.Float).Set(x).SetPrec(256).Float64()
	givenY, _ := new(big.Float).Set(y).SetPrec(256).Float64()

	response := Response{
		Action: "multiply",
		X:      givenX,
		Y:      givenY,
		Answer: answer,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}

}

func Divide(request core.Request) core.Response {
	validation, err := ValidateQueryXY(request.Query)

	if !validation {
		return core.Response{HttpStatus: http.StatusBadRequest, Data: err}
	}

	x, y := GetBigXYFromQuery(request.Query)

	answer, _ := new(big.Float).SetPrec(256).Quo(x, y).SetPrec(256).Float64()
	givenX, _ := new(big.Float).Set(x).SetPrec(256).Float64()
	givenY, _ := new(big.Float).Set(y).SetPrec(256).Float64()

	response := Response{
		Action: "divide",
		X:      givenX,
		Y:      givenY,
		Answer: answer,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}
