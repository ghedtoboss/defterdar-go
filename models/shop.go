package models

import (
	"time"
)

type Shop struct {
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Address   string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(255)"`
	OwnerID   uint      `gorm:"not_null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`

	Owner User `gorm:"foreignKey:OwnerID"`
}
