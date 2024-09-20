package models

import (
	"time"
)

type Request struct {
	ID        uint   `gorm:"primary_key"`
	ShopID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"type:uint;not null"`
	Quantity  int    `gorm:"type:int;not null"`
	Message   string `gorm:"type:text"`
	Status    string `gorm:"type:enum('pending', 'sent', 'accepted', 'rejected');default:'pending'"`
	CreatedAt time.Time

	Shop    Shop    `gorm:"foreignKey:ShopID"`
	Product Product `gorm:"foreignKey:ProductID"`
}
