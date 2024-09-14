package models

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Plan      string    `gorm:"type:enum('basic', 'premium');not null"`
	StartDate time.Time
	EndDate   time.Time
	Status    bool `gorm:"type:boolean;default:true"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
