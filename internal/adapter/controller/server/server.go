package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonsmithj/healthcheck-container-http-mysql/internal/adapter/configuration"
)

type Controller struct {
	router *gin.Engine
}

func NewController() *Controller {
	c := &Controller{
		router: gin.Default(),
	}
	c.route()
	return c
}

func (c *Controller) route() {
	r := c.router

	r.GET("/v1/healthcheck", healthCheck)
	r.GET("/v1/mysql/healthcheck", connectMysql)
	r.NoRoute(noRoute)

	r.Run(":" + configuration.Get().ListenPort)
}
