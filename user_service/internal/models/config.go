package models

type Config struct {
	Server struct {
		Host string `yaml:"host" env:"SERVER_HOST"`
		Port int    `yaml:"port" env:"SERVER_PORT"`
	} `yaml:"server"`

	Database struct {
		Username string `yaml:"username" env:"DB_USERNAME"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
		Host     string `yaml:"db_host" env:"DB_HOST"`
		Port     int    `yaml:"db_port" env:"DB_PORT"`
		DBName   string `yaml:"db_name" env:"DB_NAME"`
		SSLMode  string `yaml:"sslmode" env:"DB_SSLMODE"`
	} `yaml:"database"`
}
