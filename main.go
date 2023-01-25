package main

import "github.com/bonzzy/teltech-go-challenge/setup"

func main() {
	router := setup.SetupRouter()

	err := router.Run("localhost:8000")
	if err != nil {
		return
	}
}
