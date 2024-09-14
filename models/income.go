package models

import (
	"github.com/google/uuid"
	"time"
)

type Income struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ShopID      uuid.UUID `gorm:"type:uuid;not null"`
	Description string    `gorm:"type:text"`
	Amount      float64   `gorm:"type:float64;not null"`
	Category    string    `gorm:"type:varchar(50)"`
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Shop Shop `gorm:"foreignKey:ShopID"`
}
