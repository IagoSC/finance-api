package services

import (
	"github.com/iagosc/finance-api/api/models"
	"gorm.io/gorm"
)

type transactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) *transactionService {
	return &transactionService{
		db: db,
	}
}

func (ts *transactionService) CreateTransaction(transaction models.Transaction) {
	ts.db.Create(&transaction)
}

func (ts *transactionService) GetTransactionById(id int64) (t models.Transaction) {
	ts.db.Find(&t, id)
	return
}

func (ts *transactionService) UpdateTransaction(transaction models.Transaction) {
	ts.db.Save(&transaction)
}

func (ts *transactionService) DeleteTransaction(id int64) {
	ts.db.Delete(id)
}
