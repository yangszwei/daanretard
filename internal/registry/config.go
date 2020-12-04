package registry

import (
	"daanretard/internal/infra/config"
	"daanretard/internal/infra/validator"
)

// loadConfig load configuration from given or default path, provide
// interactive setup if no config file was found
func loadConfig(path string, v *validator.Validator, noInteractive bool) (config.Config, error) {
	var (
		cfg = config.New(v)
		err error
	)
	if path == "" {
		path = ".env"
	}
	if err = cfg.Load(path); err != nil {
		path = "./.daanretard/.env"
	}
	if err = cfg.Load(path); err != nil {
		if noInteractive {
			return config.Config{}, err
		}
		return config.InteractiveSetup()
	}
	return cfg, nil
}
