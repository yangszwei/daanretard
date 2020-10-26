package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func userLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login", gin.Params{})
}

func userRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register", gin.Params{})
}

// Setup user routes
func SetupUser(e *Engine) {
	user := e.views.Group("/user")
	user.GET("/login", userLogin)
	user.GET("/register", userRegister)
}