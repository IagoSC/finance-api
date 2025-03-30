package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	defer r.Body.Close()

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	tc.service.CreateTransaction(transaction)
}

func (tc *TransactionController) GetTransactionById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}
	transaction := tc.service.GetTransactionById(idInt)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
	return
}

func (tc *TransactionController) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	id := r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	tc.service.UpdateTransaction(idInt, transaction)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
	return
}

func (tc *TransactionController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	tc.service.DeleteTransaction(idInt)
	w.WriteHeader(http.StatusNoContent)
}
