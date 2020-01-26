package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func noRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
	return
}
