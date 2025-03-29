package services

import (
	"github.com/iagosc/finance-api/api/models"
	"gorm.io/gorm"
)

type TransactionService interface {
	CreateTransaction(transaction models.Transaction)
	GetTransaction(id int64) models.Transaction
	UpdateTransaction(transaction models.Transaction)
	DeleteTransaction(id int64)
}

type transactionService struct {
	DB *gorm.DB
}

func NewTransactionService(db *gorm.DB) *transactionService {
	return &transactionService{
		DB: db,
	}
}

func (ts *transactionService) CreateTransaction(transaction models.Transaction) {
	ret := ts.DB.Create(&transaction)
	println("Transaction created:", &ret)
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
