package configs

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig() (*conf, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	envPath := filepath.Join(rootDir, "../..", ".env")

	viper.SetConfigFile(envPath)
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, err
}
