package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iagosc/finance-api/api/controllers"
	"github.com/iagosc/finance-api/api/models"
	"github.com/iagosc/finance-api/api/routes"
	"github.com/iagosc/finance-api/api/services"
	"github.com/iagosc/finance-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initServer(".env")
}

func initServer(envPath string) {
	env := config.LoadConfig(&envPath)
	router, shutdown := create(env)
	defer shutdown()

	log.Printf("Starting server on :%s\n", env.Port)
	if err := http.ListenAndServe(":"+env.Port, router); err != nil {
		log.Fatalf("Could not start server:\n\t ERROR %s\n", err)
	}
}

func create(env *config.Env) (*http.ServeMux, func()) {
	r := http.NewServeMux()

	db, err := gorm.Open(postgres.Open(env.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database:\n\t ERROR %s\n", err)
	}

	db.AutoMigrate(&models.Transaction{}) // TODO: Remove

	routes.SetupTransactionRouter(
		r,
		controllers.NewTransactionController(
			services.NewTransactionService(db),
		),
	)
	routes.SetDefaultHandler(r)

	shutdown := func() {
		if db != nil {
			sqlDB, err := db.DB()
			if err != nil {
				log.Printf("Could not get database connection:\n\t ERROR %s\n", err)
			}
			sqlDB.Close()
		}
		log.Println("Server shutdown")
		os.Exit(0)
	}

	return r, shutdown
}
