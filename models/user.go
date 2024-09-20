package models

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:20;not null"` // e.g., admin, owner, employee etc.
	Phone     string `gorm:"type:varchar(15)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Claims struct {
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type PasswordUpdateReq struct {
	OldPassword string `form:"old_password"`
	NewPassword string `form:"new_password"`
}

type LoginReq struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
