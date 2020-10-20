// Package post define Post entity
package post

import "time"

// Post aggregate root
type Post struct {
	ID         uint32     `gorm:"autoIncrement"`
	Status     uint8      `gorm:"index"`
	Submission Submission `gorm:"foreignKey:PostID"`
	Review     Review     `gorm:"foreignKey:PostID"`
	CreatedAt  time.Time
	DeletedAt  time.Time
}