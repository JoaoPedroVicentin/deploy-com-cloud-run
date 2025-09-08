package configs

import (
	"fmt"
	"os"
)

type Config struct {
	ApiKey string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY não configurada")
	}

	return &Config{
		ApiKey: apiKey,
	}, nil
}
