package registry

import (
	"daanretard/internal/infrastructure/delivery"
	"os"
)

// NewApplication create an application
func NewApplication() *Application {
	return new(Application)
}

// Application object
type Application struct {
	engine *delivery.Engine
}

// Run start application
func (a *Application) Run() error {
	return a.engine.Run(os.Getenv("ADDR"))
}