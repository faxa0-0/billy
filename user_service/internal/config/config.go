package config

import (
	"fmt"

	"github.com/faxa0-0/billy/user_service/internal/models"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

func LoadEnv() (*models.Config, error) {
	var cfg models.Config

	err := godotenv.Load("./configs/.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file from /configs/.env: %v", err)
	}
	err = cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading environment variables: %v", err)
	}

	return &cfg, nil
}

func LoadYAML() (*models.Config, error) {
	var cfg models.Config

	err := cleanenv.ReadConfig("./configs/config.yaml", &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML config file ./configs/config.yaml: %v", err)
	}

	return &cfg, nil
}
