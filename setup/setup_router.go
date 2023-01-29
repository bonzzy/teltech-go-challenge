package setup

import (
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/core"
)

func RouterSetup() core.TinyGinServer {
	tinyGinServer := core.NewTinyGin(8000)
	tinyGinServer.Get("/healthz", controllers.Healthz)
	tinyGinServer.Get("/add", controllers.CacheWrapper(controllers.Add, controllers.Response{}))
	tinyGinServer.Get("/subtract", controllers.CacheWrapper(controllers.Subtract, controllers.Response{}))
	tinyGinServer.Get("/multiply", controllers.CacheWrapper(controllers.Multiply, controllers.Response{}))
	tinyGinServer.Get("/divide", controllers.CacheWrapper(controllers.Divide, controllers.Response{}))
	return tinyGinServer
}
