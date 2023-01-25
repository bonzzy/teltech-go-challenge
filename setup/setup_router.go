package setup

import (
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/healthz", controllers.HealthCheck)

	return router
}
