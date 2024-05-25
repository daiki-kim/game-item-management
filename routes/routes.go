package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "connected"})
	})
	return r
}
