package models

import "time"

type CashEntry struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"not null"`
	ShopID      uint `gorm:"not null"`
	CustomerID  *uint
	Amount      float64 `gorm:"not null"`
	Description string  `gorm:"size:255"`
	EntryType   string  `gorm:"size:10;not null"` // e.g., income, expense
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User     User      `gorm:"foreignKey:UserID"`
	Shop     Shop      `gorm:"foreignKey:ShopID"`
	Customer *Customer `gorm:"foreignKey:CustomerID"`
}

type GetCashEntries struct {
	FromDate time.Time `json:"from_date"`
	ToDate   time.Time `json:"to_date"`
}
