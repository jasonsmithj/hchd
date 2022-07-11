package main

import (
	"github.com/jasonsmithj/hchd/internal/adapter/configuration"
	"github.com/jasonsmithj/hchd/internal/adapter/controller/server"
)

func main() {
	initializeController()
}

func initializeController() *server.Controller {
	configuration.Load()
	controller := server.NewController()
	return controller
}
