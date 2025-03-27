package services

import (
	"github.com/iagosc/finance-api/api/models"
	"github.com/iagosc/finance-api/config"
)

type TransactionService interface {
	CreateTransaction(transaction models.Transaction)
	GetTransaction(id int64) models.Transaction
	UpdateTransaction(transaction models.Transaction)
	DeleteTransaction(id int64)
}

type transactionService struct{}

func NewTransactionService(env config.Env) *transactionService {
	return &transactionService{
		// DB: env.DB,
	}
}

func (ts *transactionService) CreateTransaction(transaction models.Transaction) {
	// ts.DB.Insert(transaction)
}

func (ts *transactionService) GetTransaction(id int64) models.Transaction {
	// return ts.DB.Find(id)
	return models.Transaction{}
}

func (ts *transactionService) UpdateTransaction(transaction models.Transaction) {
	// ts.DB.Update(transaction)
}

func (ts *transactionService) DeleteTransaction(id int64) {
	// ts.DB.Delete(id)
}
