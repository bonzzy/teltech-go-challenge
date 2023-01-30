package controllers

import (
	"github.com/bonzzy/teltech-go-challenge/core"
	"github.com/bonzzy/teltech-go-challenge/dtos"
	"github.com/bonzzy/teltech-go-challenge/helpers"
	"github.com/bonzzy/teltech-go-challenge/validations"
	"math/big"
	"net/http"
)

func Divide(request core.Request) core.Response {
	validation, err := validations.ValidateQueryXY(request.Query)

	if !validation {
		return core.Response{HttpStatus: http.StatusBadRequest, Data: err}
	}

	x, y := helpers.GetBigXYFromQuery(request.Query)

	if y.Cmp(big.NewFloat(0)) == 0 {
		return core.Response{HttpStatus: http.StatusBadRequest}
	}

	answer, _ := new(big.Float).SetPrec(256).Quo(x, y).SetPrec(256).Float64()
	givenX, _ := new(big.Float).Set(x).SetPrec(256).Float64()
	givenY, _ := new(big.Float).Set(y).SetPrec(256).Float64()

	response := dtos.Response{
		Action: "divide",
		X:      givenX,
		Y:      givenY,
		Answer: answer,
		Cached: false,
	}

	return core.Response{Data: response, HttpStatus: http.StatusOK}
}
