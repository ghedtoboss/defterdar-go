package models

import (
	"github.com/google/uuid"
	"time"
)

type Ledger struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CustomerID  uuid.UUID `gorm:"type:uuid"`
	ShopID      uuid.UUID `gorm:"type:uuid;not null"`
	Description string    `gorm:"type:varchar(255)"`
	Amount      float64   `gorm:"type:float64;not null"`
	Date        time.Time
	Status      string `gorm:"type:enum('open', 'closed');default:'open'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Customer *Customer `gorm:"foreignKey:CustomerID"`
	Shop     Shop      `gorm:"foreignKey:ShopID"`
}
