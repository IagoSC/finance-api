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

func LoadConfig(path *string) *Env {
	var envPath string
	if path != nil {
		envPath = *path
	} else {
		envPath = ".env"
	}
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading %s file", *path)
	}
	log.Println("Environment loaded from ", envPath)

	return &Env{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
	}
}
