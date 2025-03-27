package models

import (
	"database/sql"
	"time"
)

type TransactionParser struct {
	ID            int64   `json:"id"`
	Justification string  `json:"justification"`
	Type          string  `json:"type"` // "investment" | "income" | "outcome"
	Category      string  `json:"category"`
	StartDate     int64   `json:"start_date"` // epoch_day
	EndDate       int64   `json:"end_date"`   // epoch_day
	Frequency     string  `json:"frequency"`
	Amount        float64 `json:"amount"`
}

type Transaction struct {
	ID            uint
	Justification string
	Type          string
	Category      string
	StartDate     uint64
	EndDate       sql.NullInt64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
