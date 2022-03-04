package models

import (
	"time"
)

type Product struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name" gorm:"not null"`
	Desc      string         `json:"description"`
	Price     int            `json:"price" gorm:"not null"`
	Stock     int            `json:"stock"`
	Discount  int            `json:"discount"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;autoUpdateTime"`
	Thumbnail string         `json:"image_url"`
	Images    []ProductImage `json:"images"`
}

type ProductCreate struct {
	Name     string `json:"name" binding:"required"`
	Desc     string `json:"description" binding:"-"`
	Price    int    `json:"price" binding:"required"`
	Stock    int    `json:"stock" binding:"-"`
	Discount int    `json:"discount" binding:"-"`
}

type ProductUpdate struct {
	Name     string `json:"name" binding:"-"`
	Desc     string `json:"description" binding:"-"`
	Price    int    `json:"price" binding:"-"`
	Stock    int    `json:"stock" binding:"-"`
	Discount int    `json:"discount" binding:"-"`
}

type ProductImage struct {
	ID        uint `json:"id" gorm:"primary_key"`
	ProductID uint
	Order     int    `json:"order"`
	ImageUrl  string `json:"image_url" gorm:"not null"`
}

// Shipping address
