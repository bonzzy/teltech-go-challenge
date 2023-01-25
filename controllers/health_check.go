package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppHealth struct {
	Up bool `json:"up"`
}

func HealthCheck(c *gin.Context) {
	healthCheckValue := AppHealth{Up: true}
	c.JSON(http.StatusOK, healthCheckValue)
}
