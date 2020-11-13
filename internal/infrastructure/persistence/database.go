// Package persistence implement domain repositories
package persistence

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB wrap gorm.DB
type DB struct {
	Conn *gorm.DB
}

// Open connect to database
func Open(dsn string) (*DB, error) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{conn}, nil
}
