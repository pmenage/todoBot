package config

import (
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

// Config contains the API keys
type Config struct {
	TelegramKey string `json:"telegram_key"`
}

// FromYAML reads from a YAML file
func FromYAML(file []byte) Config {
	var c Config
	if err := yaml.Unmarshal(file, &c); err != nil {
		panic(errors.Wrap(err, "Could not parse YAML file"))
	}
	return c
}
