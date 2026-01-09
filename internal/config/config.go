package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Host HostConfig
	DB DBConfig
}

type HostConfig struct {
	Port	string `env:HOST_PORT,envDefault:"8080"`
}

type DBConfig struct {
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	Name     string `env:"DB_NAME,required"`
	Host     string `env:"DB_HOST,required"`
	Port 	 string `env:DP_PORT,envDefault:"5432"`
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	config := &Config{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}

	return config
}