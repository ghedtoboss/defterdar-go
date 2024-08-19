package models

import "time"

type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100"`
	Phone     string `gorm:"size:20"`
	Address   string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
