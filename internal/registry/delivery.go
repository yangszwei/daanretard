package registry

import (
	"daanretard/internal/infrastructure/application"
	"daanretard/internal/infrastructure/delivery"
	"os"
)

// PrepareDelivery prepare delivery
func PrepareDelivery(s *application.Services) (*delivery.Engine, error) {
	e := delivery.NewEngine()
	// middlewares
	delivery.SetupSession(e, s, []byte(os.Getenv("SECRET")))
	e.ApplyMiddlewares()
	// routes
	e.SetResourceRoot("./ui")
	delivery.SetupUser(e)
	delivery.SetupAPI(e)
	delivery.SetupUserAPI(e, s)
	return e, nil
}
