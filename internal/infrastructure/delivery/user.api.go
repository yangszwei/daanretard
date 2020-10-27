package delivery

import (
	"daanretard/internal/infrastructure/application"
	"daanretard/internal/object"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func apiUserRegister(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id, err := s.User.Register(object.UserProps{
			Email:      c.PostForm("email"),
			Password:   c.PostForm("password"),
			Profile:    object.UserProfileProps{
				DisplayName: c.PostForm("display_name"),
				FirstName:   c.PostForm("first_name"),
				LastName:    c.PostForm("last_name"),
			},
		})
		if err != nil {
			if err.Error() == "invalid credentials" {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid credentials",
				})
			} else if err.Error() == "email already taken" {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Email already taken",
				})
			}
			return
		}
		session.Set("user", id)
		err = session.Save()
		if err != nil {
			panic(err)
		}
		c.Status(http.StatusOK)
	}
}

func apiUserAuthenticate(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id, err := s.User.Authenticate(c.PostForm("email"), c.PostForm("password"))
		if err != nil && err.Error() == "invalid credentials" {
			c.Status(http.StatusUnauthorized)
			return
		} else if err != nil {
			panic(err)
		}
		session.Set("user", id)
		err = session.Save()
		if err != nil {
			panic(err)
		}
		c.Status(http.StatusOK)
	}
}

func apiUserUpdateEmail(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user")
		if id == nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		err := s.User.UpdateEmail(id.(uint32), c.PostForm("email"))
		if err != nil {
			switch err.Error() {
			case "record not found":
			case "invalid credentials":
				c.Status(http.StatusBadRequest)
			default:
				panic(err)
			}
			return
		}
		c.Status(http.StatusOK)
	}
}

func apiUserUpdatePassword(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user")
		if id == nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		err := s.User.UpdatePassword(id.(uint32), c.PostForm("password"))
		if err != nil {
			switch err.Error() {
			case "record not found":
			case "invalid credentials":
				c.Status(http.StatusBadRequest)
			default:
				panic(err)
			}
			return
		}
		c.Status(http.StatusOK)
	}
}

func apiUserUpdateProfile(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user")
		if id == nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		err := s.User.UpdateProfile(id.(uint32), object.UserProfileProps{
			DisplayName: c.PostForm("display_name"),
			FirstName:   c.PostForm("first_name"),
			LastName:    c.PostForm("last_name"),
		})
		if err != nil {
			switch err.Error() {
			case "record not found":
			case "invalid credentials":
				c.Status(http.StatusBadRequest)
			default:
				panic(err)
			}
			return
		}
		c.Status(http.StatusOK)
	}
}

func apiUserDelete(s *application.Services) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user")
		if id == nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		if s.User.AuthenticateWithID(id.(uint32), c.PostForm("password")) != nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		err := s.User.Delete(id.(uint32))
		if err != nil {
			if err.Error() == "record not found" {
				c.Status(http.StatusUnauthorized)
				return
			}
			panic(err)
		}
		c.Status(http.StatusOK)
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
		c.Status(http.StatusOK)
	}
}

// SetupUserAPI setup user api
func SetupUserAPI(e *Engine, s *application.Services) {
	user := e.api.Group("/user")
	user.POST("/", apiUserRegister(s))
	user.PUT("/email", apiUserUpdateEmail(s))
	user.PUT("/password", apiUserUpdatePassword(s))
	user.PUT("/profile", apiUserUpdateProfile(s))
	user.DELETE("/", apiUserDelete(s))
	user.POST("/session", apiUserAuthenticate(s))
	user.DELETE("/session", apiUserLogout())
}