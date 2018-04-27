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

	ge.GET("/doc", func(c *gin.Context) {
		location := "/doc-api"

		c.Redirect(http.StatusMovedPermanently, location)
	})

	return ge
}
