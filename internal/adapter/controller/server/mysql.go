package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonsmithj/healthcheck-container-http-mysql/internal/adapter/persistence/mysql/model"
)

func connectMysql(c *gin.Context) {
	err := model.Healthcheck()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Succeeded",
		})
	}
	return
}
