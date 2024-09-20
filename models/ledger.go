package models

import (
	"time"
)

type Ledger struct {
	ID          uint    `gorm:"primary_key"`
	CustomerID  *uint   `gorm:"type:uint"`
	ShopID      uint    `gorm:"type:uint;not null"`
	Description string  `gorm:"type:varchar(255)"`
	Amount      float64 `gorm:"type:double;not null"`
	Date        time.Time
	Status      string `gorm:"type:enum('open', 'closed');default:'open'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Customer *Customer `gorm:"foreignKey:CustomerID"`
	Shop     Shop      `gorm:"foreignKey:ShopID"`
}
