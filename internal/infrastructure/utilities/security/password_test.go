package security_test

import (
	"daanretard/internal/infrastructure/utilities/security"
	"testing"
)

var (
	hash     []byte
	password = "12345678"
)

func TestGenerateFromPassword(t *testing.T) {
	var err error
	hash, err = security.GenerateFromPassword(password)
	if err != nil {
		t.Error(err)
	}
}

func TestCompareHashAndPassword(t *testing.T) {
	err := security.CompareHashAndPassword(hash, password)
	if err != nil {
		t.Error(err)
	}
}
