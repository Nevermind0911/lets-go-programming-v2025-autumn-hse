package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	SourceFile string `yaml:"input-file"`
	TargetFile string `yaml:"output-file"`
}

func Load(path string) (*AppConfig, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var cfg AppConfig
	if err := yaml.Unmarshal(bytes, &cfg); err != nil {
		return nil, fmt.Errorf("invalid YAML format: %w", err)
	}

	return &cfg, nil
}
