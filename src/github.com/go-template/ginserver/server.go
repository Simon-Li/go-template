package ginserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Engine create a gin engine
func Engine() *gin.Engine {
	ge := gin.New()

	ge.Use(gin.Recovery())

	ge.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return ge
}
