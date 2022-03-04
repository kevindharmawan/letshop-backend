package models

import (
	"time"
)

// TODO: Add address and bank account
type Shop struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	UserID    string    `json:"user_id" gorm:"type:varchar(30);unique;not null"`
	Name      string    `json:"name" gorm:"not null"`
	Desc      string    `json:"description"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;autoUpdateTime"`
	ImageUrl  string    `json:"image_url"`
	BannerUrl string    `json:"banner_url"`
	Product   []Product `json:"products"`
}

type ShopCreate struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"description" binding:"-"`
}

type ShopUpdate struct {
	Name string `json:"name" binding:"-"`
	Desc string `json:"description" binding:"-"`
}
