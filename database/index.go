package main

import (
	"log"
	"os"

	"github.com/iagosc/finance-api/api/models"
	"github.com/iagosc/finance-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cmd := "testConn"

	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	env := config.LoadConfig(nil)

	// Initialize the database connection
	db, err := gorm.Open(postgres.Open(env.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Panic("Failed to connect to database")
	}

	switch cmd {
	case "migrate":
		log.Print("Running migrations...")
		err = db.AutoMigrate(&models.Transaction{})
		if err != nil {
			log.Panic("Failed to migrate database")
		}
	case "testConn":
	default:
		log.Print("Testing connections...")
		DB, err := db.DB()
		if err != nil {
			log.Panic("Failed to get database connection")
		}
		if err := DB.Ping(); err != nil {
			log.Panic("Failed to ping database")
		}
	}
}
