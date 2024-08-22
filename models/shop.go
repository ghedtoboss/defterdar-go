package models

import "time"

type Shop struct {
	ID        uint   `gorm:"primaryKey"`
	OwnerID   uint   `gorm:"not null"`
	Name      string `gorm:"size:100;not null"`
	Address   string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Owner User `gorm:"foreignKey:OwnerID"`
}
