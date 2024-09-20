package models

import (
	"time"
)

type Sale struct {
	ID          uint    `gorm:"primary_key"`
	ShopID      uint    `gorm:"type:uint;not null"`
	CustomerID  *uint   `gorm:"type:uint"`
	TotalAmount float64 `gorm:"type:double;not null"`
	Discount    float64 `gorm:"type:double"`
	CreatedAt   time.Time

	Shop     Shop      `gorm:"foreignKey:ShopID"`
	Customer *Customer `gorm:"foreignKey:CustomerID"`
}
