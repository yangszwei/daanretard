package delivery_test

import (
	"daanretard/internal/infra/delivery"
	"testing"
)

func createTestServer() *delivery.Server {
	s := delivery.NewServer("../../../ui")
	delivery.SetupSession(s, []byte("secret"))
	s.LoadTemplates()
	s.SetupRouterGroups()
	return s
}

func TestNewServer(t *testing.T) {
	_ = delivery.NewServer("")
}

func TestServer_LoadTemplates(t *testing.T) {
	delivery.NewServer("").LoadTemplates()
}

func TestServer_SetupRouters(t *testing.T) {
	delivery.NewServer("").SetupRouterGroups()
}
