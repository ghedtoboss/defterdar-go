package models

import "time"

type Invoice struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `gorm:"not null"`
	CustomerID uint    `gorm:"not null"`
	Amount     float64 `gorm:"not null"`
	Status     string  `gorm:"size:20;not null"` // e.g., paid, unpaid
	IssueDate  time.Time
	DueDate    time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	User     User     `gorm:"foreignKey:UserID"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
}
