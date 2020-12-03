package config

import (
	"daanretard/internal/infra/errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Load load config from .env file
func Load(path string) (Config, error) {
	if path == "" {
		path = ".env"
	}
	if err := godotenv.Load(path); err != nil {
		return Config{}, errors.New(fmt.Sprintf("No .env file found at: %s", path))
	}
	return Config{
		Addr:        os.Getenv("ADDR"),
		Secret:      os.Getenv("SECRET"),
		FbPageID:    os.Getenv("FB_PAGE_ID"),
		FbAppID:     os.Getenv("FB_APP_ID"),
		FbAppSecret: os.Getenv("FB_APP_SECRET"),
		Data:        os.Getenv("DATA"),
		DbDsn:       os.Getenv("DB_DSN"),
	}, nil
}

// Config object
type Config struct {
	Addr        string
	Secret      string
	FbPageID    string
	FbAppID     string
	FbAppSecret string
	Data        string
	DbDsn       string
}
