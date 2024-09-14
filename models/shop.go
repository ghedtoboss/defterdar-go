package models

import (
	"github.com/google/uuid"
	"time"
)

type Shop struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Address   string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255)"`
	OwnerID   uuid.UUID `gorm:"type:uuid;unique;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Owner *User `gorm:"foreignKey:OwnerID"`
}
