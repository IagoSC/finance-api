package routes

import (
	"net/http"

	"github.com/iagosc/finance-api/api/controllers"
)

func SetupTransactionRouter(transactionController *controllers.TransactionController) *http.ServeMux {
	transactionSubRouter := http.NewServeMux()
	transactionSubRouter.HandleFunc("POST /", transactionController.CreateTransaction)
	transactionSubRouter.HandleFunc("GET /{id}", transactionController.GetTransactionById)
	transactionSubRouter.HandleFunc("PUT /{id}", transactionController.UpdateTransaction)
	transactionSubRouter.HandleFunc("DELETE /{id}", transactionController.DeleteTransaction)
	return transactionSubRouter
}
