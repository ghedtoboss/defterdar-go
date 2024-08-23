package models

import "time"

type Employee struct {
	ID        uint   `gorm:"primaryKey"`
	ShopID    uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Role      string `gorm:"not null"` //e.g.
	CreatedAt time.Time
	UpdatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
	Shop Shop `gorm:"foreignKey:ShopID"`
}

type EmployeeRoleUpdate struct {
	Role string `json:"role"`
}
