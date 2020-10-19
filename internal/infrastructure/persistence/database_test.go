package persistence_test

import (
	"daanretard/internal/infrastructure/persistence"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

var DB *persistence.DB

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic(err)
	}
}

func TestOpen(t *testing.T) {
	var err error
	DB, err = persistence.Open(os.Getenv("DB_DSN"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpen_Error(t *testing.T) {
	var err error
	_, err = persistence.Open("invalid dsn")
	if err == nil {
		t.Fatal("Expected error but none present")
	}
}
