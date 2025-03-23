package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	DATABASE_URL string
}
type AuthConfig struct {
	AccessToken  string
	RefreshToken string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Db: DbConfig{
			DATABASE_URL: os.Getenv("DATABASE_URL"),
		},
		Auth: AuthConfig{
			AccessToken:  os.Getenv("ACCESS_TOKEN"),
			RefreshToken: os.Getenv("REFRESH_TOKEN"),
		},
	}
}
