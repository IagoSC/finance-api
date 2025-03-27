package main

import (
	"log"
	"net/http"
	"os"

	"github.com/iagosc/finance-api/api/controllers"
	"github.com/iagosc/finance-api/api/routes"
	"github.com/iagosc/finance-api/api/services"
	"github.com/iagosc/finance-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	server(os.Args[1])
	log.Println("Starting server on :8080")
}

func server(envPath string) {
	env := config.LoadConfig(envPath)
	router, shutdown := create(env)
	defer shutdown()

	if err := http.ListenAndServe(env.Port, router); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func create(env *config.Env) (*http.ServeMux, func()) {
	r := http.NewServeMux()

	gorm.Open(postgres.Open(postgres.Config{
		DSN: env.DatabaseURL,
	}), &gorm.Config{})
	routes.SetupTransactionRouter(r, controllers.NewTransactionController(services.NewTransactionService()))

	shutdown := func() {}

	return r, shutdown
}
