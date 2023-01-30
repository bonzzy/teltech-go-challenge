package controllers

import (
	"github.com/bonzzy/teltech-go-challenge/core"
	"github.com/bonzzy/teltech-go-challenge/dtos"
	"github.com/bonzzy/teltech-go-challenge/helpers"
	"github.com/bonzzy/teltech-go-challenge/validations"
	"math/big"
	"net/http"
)

func Multiply(request core.Request) core.Response {
	validation, err := validations.ValidateQueryXY(request.Query)

	if !validation {
		return core.Response{HttpStatus: http.StatusBadRequest, Data: err}
	}

	x, y := helpers.GetBigXYFromQuery(request.Query)

	answer, _ := new(big.Float).SetPrec(256).Mul(x, y).SetPrec(256).Float64()
	givenX, _ := new(big.Float).Set(x).SetPrec(256).Float64()
	givenY, _ := new(big.Float).Set(y).SetPrec(256).Float64()

	response := dtos.Response{
		Action: "multiply",
		X:      givenX,
		Y:      givenY,
		Answer: answer,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}
