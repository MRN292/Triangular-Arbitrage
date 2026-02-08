package config

import (
	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
)

type Config struct {
	APIConfig *API
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()
	apiConfig := new(API)
	if err := env.Set(apiConfig); err != nil {
		return nil, err
	}

	config := Config{
		APIConfig: apiConfig,
	}

	return &config, nil
}
