package main

import (
	"github.com/bonzzy/teltech-go-challenge/setup"
)

func main() {
	tinyGinServer := setup.RouterSetup()
	tinyGinServer.Run()
}
