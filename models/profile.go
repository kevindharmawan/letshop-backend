package models

import (
	"time"
)

type Profile struct {
	ID                string            `json:"user_id" gorm:"type:varchar(30);unique;primary_key"`
	Name              string            `json:"name" gorm:"not null"`
	BirthDate         time.Time         `json:"birth_date" gorm:"not null"`
	CreatedAt         time.Time         `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt         time.Time         `json:"updated_at" gorm:"not null;autoUpdateTime"`
	ImageUrl          string            `json:"image_url"`
	Gender            string            `json:"gender" gorm:"not null"`
	ShippingAddresses []ShippingAddress `json:"shipping_addresses" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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

type ShippingAddress struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserID    string `json:"-"`
	Name      string `json:"name" gorm:"not null"`
	Recipient string `json:"recipient" gorm:"not null"`
	Phone     int    `json:"phone" gorm:"not null"`
	Address   string `json:"address" gorm:"not null"`
}

type ShippingAddressCreate struct {
	Name      string `json:"name" binding:"required"`
	Recipient string `json:"recipient" binding:"required"`
	Phone     int    `json:"phone" binding:"required"`
	Address   string `json:"address" binding:"required"`
}

type ShippingAddressUpdate struct {
	Name      string `json:"name" binding:"-"`
	Recipient string `json:"recipient" binding:"-"`
	Phone     int    `json:"phone" binding:"-"`
	Address   string `json:"address" binding:"-"`
}
