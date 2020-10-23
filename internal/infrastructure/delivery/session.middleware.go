package delivery

import (
	"daanretard/internal/infrastructure/application"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// SetupSession set session middleware
func SetupSession(e *Engine, s *application.Services, secret []byte) {
	store := cookie.NewStore(secret)
	e.engine.Use(sessions.Sessions("session", store))
}