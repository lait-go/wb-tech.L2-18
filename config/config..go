package config

import (
	"calendar/pkg/httpserver"
	"calendar/pkg/logger"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Http   httpserver.Config `yaml:"http"`
	Logger logger.Config     `yaml:"logger"`
}

func New() (Config, error) {
	file, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return Config{}, fmt.Errorf("Config error: %s", err.Error())
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return Config{}, fmt.Errorf("Config error: %s", err.Error())
	}

	return cfg, nil
}
