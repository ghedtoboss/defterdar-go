package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"size:100;unique;not null"`
	Password  string    `gorm:"size:255;not null"`
	Role      string    `gorm:"size:20;not null"` // e.g., admin, owner, employee etc.
	Phone     string    `gorm:"type:varchar(15)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Claim struct {
	Email  string `json:"email"`
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type PasswordUpdateReq struct {
	OldPassword string `form:"old_password"`
	NewPassword string `form:"new_password"`
}

type LoginReq struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
