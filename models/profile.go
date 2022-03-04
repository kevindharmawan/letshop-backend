package models

import (
	"time"
)

type Profile struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    string    `json:"user_id" gorm:"type:varchar(30);unique;not null"`
	Name      string    `json:"name" gorm:"not null"`
	BirthDate time.Time `json:"birth_date" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
	ImageUrl  string    `json:"image_url"`
	Gender    string    `json:"gender" gorm:"not null"`
}

type ProfileCreate struct {
	Name      string    `json:"name" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
	Gender    string    `json:"gender" binding:"required"`
}

type ProfileUpdate struct {
	Name      string    `json:"name" binding:"-"`
	BirthDate time.Time `json:"birth_date" binding:"-"`
	Gender    string    `json:"gender" binding:"-"`
}

// Shipping address
