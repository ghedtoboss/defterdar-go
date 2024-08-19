package models

import "time"

type CustomerTransaction struct {
	ID              uint    `gorm:"primaryKey"`
	CustomerID      uint    `gorm:"not null"`
	Amount          float64 `gorm:"not null"`
	TransactionType string  `gorm:"size:10;not null"` // e.g., credit, debit
	Description     string  `gorm:"size:255"`
	CreatedAt       time.Time
	UpdatedAt       time.Time

	Customer Customer `gorm:"foreignKey:CustomerID"`
}
