package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DatabaseURL string
	Port        string
}

func LoadConfig(path string) *Env {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Env{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}
