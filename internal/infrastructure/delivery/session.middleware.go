package delivery

import (
	"daanretard/internal/infrastructure/application"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SetupSession set session middleware
func SetupSession(e *Engine, s *application.Services, secret []byte) {
	store := cookie.NewStore(secret)
	e.engine.Use(sessions.Sessions("session", store))
	e.engine.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		c.Set("user", session.Get("user"))
	})
}
