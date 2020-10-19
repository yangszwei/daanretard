package registry

import "daanretard/internal/infrastructure/delivery"

// PrepareDelivery prepare delivery
func PrepareDelivery(s *Services) (*delivery.Engine, error) {
	e := delivery.NewEngine()
	delivery.SetupAPI(e)
	return e, nil
}