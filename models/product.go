package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Description   string    `gorm:"type:text"`
	Price         float64   `gorm:"type:decimal(10,2);not null"`
	Cost          float64   `gorm:"type:decimal(10,2)"`
	StockQuantity int       `gorm:"type:int;not null"`
	ShopID        uuid.UUID `gorm:"type:uuid;not null"`
	SupplierName  string    `gorm:"type:varchar(100)"` // Tedarikçi adı
	SupplierEmail string    `gorm:"type:varchar(100)"` // Tedarikçi e-posta
	SupplierPhone string    `gorm:"type:varchar(15)"`  // Tedarikçi telefon
	PhotoURL      string    `gorm:"type:varchar(255)"` // Ürün fotoğrafının URL'si
	CreatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`
	UpdatedAt     time.Time `gorm:"type:timestamp;default:current_timestamp"`

	Shop Shop `gorm:"foreignKey:ShopID"`
}
