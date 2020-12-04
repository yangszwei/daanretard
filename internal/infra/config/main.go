package config

import (
	"daanretard/internal/infra/errors"
	"daanretard/internal/infra/validator"
	"encoding/json"
	"github.com/joho/godotenv"
)

// New create an instance of Config
func New(v *validator.Validator) Config {
	var c Config
	c.v = v
	return c
}

// Config object
type Config struct {
	Addr             string `json:"ADDR" validate:"required"`
	Secret           string `json:"SECRET" validate:"required"`
	FbPageID         string `json:"FB_PAGE_ID" validate:"required"`
	FbGraphAppID     string `json:"FB_GRAPH_APP_ID" validate:"required"`
	FbGraphAppSecret string `json:"FB_GRAPH_APP_SECRET" validate:"required"`
	DataPath         string `json:"DATA_PATH" validate:"required"`
	DbDsn            string `json:"DB_DSN" validate:"required"`
	Path             string `json:"-"`
	v                *validator.Validator
}

// Load load config from env file
func (c *Config) Load(path string) error {
	cfg, err := godotenv.Read(path)
	if err != nil {
		return errors.From(err)
	}
	text, _ := json.Marshal(cfg)
	_ = json.Unmarshal(text, c)
	c.Path = path
	return nil
}

// Validate validate configuration
func (c Config) Validate() error {
	return c.v.Validate(c)
}

// Save write current config to .env file, used only on first run if the
// configuration is not set
func (c Config) Save() error {
	var m map[string]string
	text, _ := json.Marshal(c)
	_ = json.Unmarshal(text, &m)
	return godotenv.Write(m, c.Path)
}
