package models

import (
	"time"
)

type Transaction struct {
	ID          uint    `gorm:"primary_key"`
	ShopID      uint    `gorm:"not null"`
	Description *string `gorm:"type:text"`
	Amount      float64 `gorm:"type:double;not null"`
	Category    *string `gorm:"type:varchar(50)"`
	Type        string  `gorm:"type:varchar(50)"` //"expense", "income"
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Shop Shop `gorm:"foreignKey:ShopID"`
}
