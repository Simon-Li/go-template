package gin_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinEngine() *gin.Engine {
	ge := gin.New()

	ge.Use(gin.Recovery())

	ge.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return ge
}
