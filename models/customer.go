package models

import (
	"github.com/google/uuid"
	"time"
)

type Customer struct {
	ID        uuid.UUID `gorm:"type:uuid;default;uuid_generate_v4();primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255)"`
	Email     string    `gorm:"type:varchar(255)"`
	ShopID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Shop Shop `gorm:"foreignKey:ShopID"`
}
