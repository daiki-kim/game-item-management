package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ItemRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/item", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "got item"})
	})
	return r
}
