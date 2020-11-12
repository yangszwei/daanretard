// Package security implement methods related to password hashing
package security

import "golang.org/x/crypto/bcrypt"

// GenerateFromPassword generate hash from password
func GenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CompareHashAndPassword compare hash and password
func CompareHashAndPassword(hash []byte, password string) error {
	return bcrypt.CompareHashAndPassword(hash, []byte(password))
}
