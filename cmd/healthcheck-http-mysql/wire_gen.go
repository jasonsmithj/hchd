package main

import (
	"github.com/jasonsmithj/healthcheck-container-http-mysql/internal/adapter/configuration"
	"github.com/jasonsmithj/healthcheck-container-http-mysql/internal/adapter/controller/server"
)

func initializeController() *server.Controller {
	configuration.Load()
	controller := server.NewController()
	return controller
}
