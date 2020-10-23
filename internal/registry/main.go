package registry

import (
	"daanretard/internal/infrastructure/persistence"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

// Prepare prepare app
func Prepare() *Application {
	app := NewApplication()
	db, err := persistence.Open(os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}
	services, err := SetupService(db)
	if err != nil {
		panic(err)
	}
	engine, err := PrepareDelivery(services)
	if err != nil {
		panic(err)
	}
	app.engine = engine
	return app
}