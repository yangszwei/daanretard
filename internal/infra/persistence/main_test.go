package persistence_test

import (
	"daanretard/internal/infra/persistence"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func init() {
	if err := godotenv.Load("../../../.env"); err != nil {
		panic(err)
	}
}

// newDB create a DB
func newDB() *persistence.DB {
	db, err := persistence.Open(os.Getenv("DB_DSN"))
	if err != nil {
		panic(err)
	}
	return db
}

func TestDB(t *testing.T) {
	if err := newDB().Close(); err != nil {
		t.Error(err)
	}
}
