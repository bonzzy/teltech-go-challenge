package setup

import (
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/core"
	"github.com/bonzzy/teltech-go-challenge/dtos"
	"github.com/bonzzy/teltech-go-challenge/middleware"
	"os"
	"strconv"
)

func RouterSetup() core.TinyGinServer {
	portEnv := os.Getenv("PORT")
	port, _ := strconv.Atoi(portEnv)

	if port == 0 {
		port = 8080
	}

	tinyGinServer := core.NewTinyGin(port)
	tinyGinServer.Get("/healthz", controllers.Healthz)
	tinyGinServer.Get("/add", middleware.CacheWrapper(controllers.Add, dtos.Response{}))
	tinyGinServer.Get("/subtract", middleware.CacheWrapper(controllers.Subtract, dtos.Response{}))
	tinyGinServer.Get("/multiply", middleware.CacheWrapper(controllers.Multiply, dtos.Response{}))
	tinyGinServer.Get("/divide", middleware.CacheWrapper(controllers.Divide, dtos.Response{}))

	return tinyGinServer
}
