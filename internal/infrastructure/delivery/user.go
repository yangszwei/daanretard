package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func userHome(c *gin.Context) {
	user, _ := c.Get("user")
	c.HTML(http.StatusOK, "user/home", gin.H{
		"user": user,
	})
}

func userLogin(c *gin.Context) {
	user, _ := c.Get("user")
	if user != nil {
		c.Redirect(http.StatusFound, "/user/home")
		return
	}
	c.HTML(http.StatusOK, "user/login", gin.H{
		"user": false,
	})
}

func userRegister(c *gin.Context) {
	user, _ := c.Get("user")
	if user != nil {
		c.Redirect(http.StatusFound, "/user/home")
		return
	}
	c.HTML(http.StatusOK, "user/register", gin.H{
		"user": false,
	})
}

// Setup user routes
func SetupUser(e *Engine) {
	user := e.views.Group("/user")
	user.GET("/home", userHome)
	user.GET("/login", userLogin)
	user.GET("/register", userRegister)
}