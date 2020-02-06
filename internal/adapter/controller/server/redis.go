package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonsmithj/hchd/internal/adapter/persistence/redis/model"
)

func connectRedis(c *gin.Context) {
	err := model.RedisHealthcheck()
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
