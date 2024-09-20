package models

import (
	"time"
)

type Customer struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Phone     string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(255)"`
	ShopID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Shop Shop `gorm:"foreignKey:ShopID"`
}
