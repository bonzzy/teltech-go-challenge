package setup

import (
	"github.com/bonzzy/teltech-go-challenge/controllers"
	"github.com/bonzzy/teltech-go-challenge/core"
)

func RouterSetup() core.TinyGinServer {
	tinyGinServer := core.NewTinyGin(8000)
	tinyGinServer.Get("/healthz", controllers.Healthz)
	tinyGinServer.Get("/add", controllers.Add)
	tinyGinServer.Get("/substract", controllers.Subtract)
	tinyGinServer.Get("/multiply", controllers.Multiply)
	tinyGinServer.Get("/divide", controllers.Divide)

	tinyGinServer.Run()
	return tinyGinServer
}
