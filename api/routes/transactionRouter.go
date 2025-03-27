package routes

import (
	"net/http"

	"github.com/iagosc/finance-api/api/controllers"
)

func SetupTransactionRouter(router *http.ServeMux, transactionController *controllers.TransactionController) {
	router.HandleFunc("POST /transactions", transactionController.CreateTransaction)
	// router.HandleFunc("GET /transactions/:id", transactionController.GetTransaction)
	// router.HandleFunc("PUT /transactions/:id", transactionController.UpdateTransaction)
	// router.HandleFunc("DELETE /transactions/:id", transactionController.DeleteTransaction)
}
