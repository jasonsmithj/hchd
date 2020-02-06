package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonsmithj/hchd/internal/adapter/persistence/postgres/model"
)

func connectPostgres(c *gin.Context) {
	err := model.PostgresHealthcheck()
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
