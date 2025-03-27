package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/iagosc/finance-api/api/models"
	"github.com/iagosc/finance-api/api/services"
)

type TransactionController struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (tc *TransactionController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	tc.service.CreateTransaction(transaction)

}
