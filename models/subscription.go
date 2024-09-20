package models

import (
	"time"
)

type Subscription struct {
	ID        uint   `gorm:"primary_key"`
	UserID    uint   `gorm:"type:uint;not null"`
	Plan      string `gorm:"type:enum('basic', 'premium');not null"`
	StartDate time.Time
	EndDate   time.Time
	Status    bool `gorm:"type:boolean;default:true"`
	CreatedAt time.Time

	User User `gorm:"foreignKey:UserID"`
}
