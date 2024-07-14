package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBDriver   string
	Port       string
	DBUri      string
}

func NewConfig() (*Config, error) {
	envMap, err := godotenv.Read()
	if err != nil {
		return nil, err
	}
	return &Config{
		DBUser:     envMap["DB_USER"],
		DBPassword: envMap["DB_PASSWORD"],
		DBName:     envMap["DB_NAME"],
		DBPort:     envMap["DB_PORT"],
		DBDriver:   envMap["DB_DRIVER"],
		Port:       envMap["PORT"],
		DBUri:      envMap["DB_URI"],
	}, nil
}
