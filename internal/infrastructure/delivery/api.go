package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupAPI setup api
func SetupAPI(e *Engine) {
	e.api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version": "0.0.1",
		})
	})
}
