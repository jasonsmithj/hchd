package main

import (
	"github.com/jasonsmithj/hchd/internal/adapter/configuration"
	"github.com/jasonsmithj/hchd/internal/adapter/controller/server"
)

func initializeController() *server.Controller {
	configuration.Load()
	controller := server.NewController()
	return controller
}
