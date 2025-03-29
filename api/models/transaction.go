package models

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID            uint
	Amount        float64
	Justification string
	Type          string
	Category      string
	StartDate     uint64
	EndDate       sql.NullInt64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
