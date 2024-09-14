package models

import (
	"github.com/google/uuid"
	"time"
)

type Request struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	ShopID    uuid.UUID `gorm:"type:uuid;not null"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	Quantity  int       `gorm:"type:int;not null"`
	Message   string    `gorm:"type:text"`
	Status    string    `gorm:"type:enum('pending', 'sent', 'accepted', 'rejected');default:'pending'"`
	CreatedAt time.Time

	Shop    Shop    `gorm:"foreignKey:ShopID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
