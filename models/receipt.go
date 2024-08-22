package models

import "time"

type Receipt struct {
	ID          uint    `gorm:"primaryKey"`
	UserID      uint    `gorm:"not null"`
	ShopID      uint    `gorm:"not null"`
	CustomerID  uint    `gorm:"not null"`
	Amount      float64 `gorm:"not null"`
	Description string  `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	User     User     `gorm:"foreignKey:UserID"`
	Customer Customer `gorm:"foreignKey:CustomerID"`
	Shop     Shop     `gorm:"foreignKey:ShopID"`
}
