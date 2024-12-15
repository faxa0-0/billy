package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	DBPath string

	Address string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &Config{}

	config.Address = cast.ToString(getEnvValue("HTTP_SERVER_ADDRESS", "localhost:8881"))
	config.DBPath = cast.ToString(getEnvValue("DB_PATH", "./db.db"))

	return config, nil
}
func getEnvValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultValue
}
