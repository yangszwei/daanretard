//+build !github

package config_test

import (
	"daanretard/internal/infra/config"
	"testing"
)

func TestInteractiveSetup(t *testing.T) {
	_, err := config.InteractiveSetup()
	if err != nil {
		t.Error(err)
	}
}
