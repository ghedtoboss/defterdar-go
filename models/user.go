package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:20;not null"` // e.g., admin, user, etc.
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PasswordUpdateReq struct {
	OldPassword string `form:"old_password"`
	NewPassword string `form:"new_password"`
}

type LoginReq struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
