package controllers

import (
	"github.com/bonzzy/teltech-go-challenge/core"
	"net/http"
)

type AppHealth struct {
	Up bool `json:"up"`
}

func Healthz(_ core.Request) core.Response {
	return core.Response{Data: AppHealth{Up: true}, HttpStatus: http.StatusOK}
}
