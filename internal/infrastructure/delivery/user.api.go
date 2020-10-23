package delivery

import (
	"daanretard/internal/infrastructure/application"
	"daanretard/internal/object"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiUserRegister(services *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID, err := services.User.Register(
			object.UserProps{
				Name: c.PostForm("name"),
				Email: c.PostForm("email"),
				Password: c.PostForm("password"),
			},
			object.UserProfileProps{
				FirstName: c.PostForm("first_name"),
				LastName:  c.PostForm("last_name"),
			},
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		session.Set("user", userID)
		err = session.Save()
		if err != nil {
			panic(err)
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

func apiUserLogin(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id, err := s.User.Authenticate(object.UserProps{
			Name:     c.PostForm("name"),
			Email:    c.PostForm("email"),
			Password: c.PostForm("password"),
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		session.Set("user", id)
		err = session.Save()
		if err != nil {
			panic(err)
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

func apiUserLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("user")
		err := session.Save()
		if err != nil {
			panic(err)
		}
		c.AbortWithStatus(http.StatusOK)
	}
}

// SetupUserAPI setup user api
func SetupUserAPI(e *Engine, s *application.Services) {
	user := e.api.Group("/user")
	user.POST("/", apiUserRegister(s))
	user.POST("/session", apiUserLogin(s))
	user.DELETE("/session", apiUserLogout())
}