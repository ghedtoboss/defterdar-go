package models

import "time"

type CashEntry struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	ShopID      uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Description string  `gorm:"size:255"`
	EntryType   string  `gorm:"size:10;not null"` // e.g., income, expense
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User User `gorm:"foreignKey:UserID"`
	Shop Shop `gorm:"foreignKey:ShopID"`
}
