package delivery

import (
	"daanretard/internal/entity/user"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func getQuery(key string, c *gin.Context) string {
	query := c.Request.URL.Query()
	if _, e := query[key]; e {
		return query[key][0]
	}
	return ""
}

// apiUserAuth user register / sign in with facebook access token
func apiUserAuth(u user.IUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id, err := u.ContinueWithFacebook(c.PostForm("access_token"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		session.Set("user", id)
		if err := session.Save(); err != nil {
			panic(err)
		}
	}
}

// apiUserGetProfile return user profile with query: id, name, expires_at, created_at
func apiUserGetProfile(u user.IUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user")
		if userID == nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		record, err := u.GetProfile(userID.(string))
		if err != nil && err.Error() == "record not found" {
			c.Status(http.StatusUnauthorized)
			return
		} else if err != nil {
			panic(err)
		}
		query := getQuery("fields", c)
		if query == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no field selected",
			})
		}
		result := make(map[string]interface{})
		fields := strings.Split(query, ",")
		// TODO: optimize this method
		for _, field := range fields {
			if field == "id" {
				result["id"] = userID
			} else if field == "name" {
				result["name"] = record.Name
			} else if field == "expires_at" {
				result["is_verified"] = record.ExpiresAt
			} else if field == "created_at" {
				result["created_at"] = record.CreatedAt
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		})
	}
}

// apiUserDelete delete user (require access token)
func apiUserDelete(u user.IUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user")
		fmt.Println(session.Get("user"))
		if userID == nil {
			c.Status(http.StatusUnauthorized)
			return
		}
		// TODO: create a method for verifying user without creating a new one if not exist
		id, err := u.ContinueWithFacebook(c.PostForm("access_token"))
		if err != nil || id != userID {
			c.Status(http.StatusUnauthorized)
			return
		}
		if err := u.Delete(id); err != nil {
			panic(err)
		}
		session.Delete("user")
		if err := session.Save(); err != nil {
			panic(err)
		}
		c.Status(http.StatusOK)
		return
	}
}

// apiUserSignOut sign user out
func apiUserSignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete("user")
		if err := session.Save(); err != nil {
			panic(err)
		}
	}
}

// SetupUserAPI add user api routes to server
func SetupUserAPI(s *Server, u user.IUsecase) {
	g := s.api.Group("/user")
	g.POST("/auth", apiUserAuth(u))
	g.DELETE("/auth", apiUserSignOut())
	g.GET("/", apiUserGetProfile(u))
	g.POST("/delete", apiUserDelete(u))
}
