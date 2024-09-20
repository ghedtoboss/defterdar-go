package models

import (
	"time"
)

type Product struct {
	ID            uint      `gorm:"primary_key"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Description   string    `gorm:"type:text"`
	Price         float64   `gorm:"type:double;not null"`
	Cost          float64   `gorm:"type:double"`
	StockQuantity int       `gorm:"type:int;not null"`
	ShopID        uint      `gorm:"type:uint;not null"`
	SupplierName  string    `gorm:"type:varchar(100)"` // Tedarikçi adı
	SupplierEmail string    `gorm:"type:varchar(100)"` // Tedarikçi e-posta
	SupplierPhone string    `gorm:"type:varchar(15)"`  // Tedarikçi telefon
	PhotoURL      string    `gorm:"type:varchar(255)"` // Ürün fotoğrafının URL'si
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`

	Shop Shop `gorm:"foreignKey:ShopID"`
}
