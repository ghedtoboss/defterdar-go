package models

import (
	"github.com/google/uuid"
	"time"
)

type Sale struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ShopID      uuid.UUID `gorm:"type:uuid;not null"`
	CustomerID  uuid.UUID `gorm:"type:uuid"`
	TotalAmount float64   `gorm:"type:float;not null"`
	Discount    float64   `gorm:"type:float"`
	CreatedAt   time.Time

	Shop     Shop      `gorm:"foreignKey:ShopID"`
	Customer *Customer `gorm:"foreignKey:CustomerID"`
}
