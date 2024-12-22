package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        string    `gorm:"type:char(36);primary_key" json:"id,omitempty"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex:idx_users_email,LENGTH(255);not null" json:"email,omitempty"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	CreatedAt time.Time `gorm:"not null;default:'1970-01-01 00:00:01'" json:"createdAt,omitempty"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New().String()
	return nil
}

type CreateUserSchema struct {
	Name     string `json:"name,omitempty" validate:"required,min=3,max=255"`
	Email    string `json:"email,omitempty" validate:"required,email,min=3,max=255"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=255"`
}

type UpdateUserSchema struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=3,max=255"`
	Email    string `json:"email,omitempty" validate:"omitempty,email,min=3,max=255"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6,max=255"`
}

type UserResponseSchema struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
