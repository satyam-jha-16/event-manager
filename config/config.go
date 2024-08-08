package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT, required"`
	DBHost     string `env:"DB_HOST, required"`
	DBName     string `env:"DB_NAME, required"`
	DBUser     string `env:"DB_USER, required"`
	DBPassword string `env:"DB_PASSWORD, required"`
	DBSSLMODE  string `env:"DB_SSL_MODE, required"`
}


func NewEnvConfig() *EnvConfig {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file %e",err )
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Error parsing env config %e",err )
	}
	return config
}
