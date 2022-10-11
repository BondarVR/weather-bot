package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Token       string `env:"TELEGRAM_TOKEN"`
	LogLevel    string `env:"LOG_LEVEL"`
	LogServer   string `env:"LOG_SERVER"`
	ServiceName string `env:"SERVICE_NAME"`
	ApiWeather  string `env:"API_WEATHER"`
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var config Config

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, err
}
