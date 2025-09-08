package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig() (*Config, error) {
	var config Config

	rootDir, err := os.Getwd()
	if err == nil {
		envPath := filepath.Join(rootDir, "../..", ".env")
		viper.SetConfigFile(envPath)
		_ = viper.ReadInConfig()
	}

	viper.AutomaticEnv()

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if config.ApiKey == "" {
		config.ApiKey = os.Getenv("API_KEY")
	}

	if config.ApiKey == "" {
		return nil, fmt.Errorf("API_KEY não configurada")
	}

	return &config, nil
}
