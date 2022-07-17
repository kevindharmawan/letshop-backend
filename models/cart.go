package models

type Cart struct {
	ID             uint    `json:"id" gorm:"primary_key"`
	UserID         string  `json:"-"`
	User           Profile `json:"-" gorm:"foreignKey:UserID"`
	ProductID      uint    `json:"-"`
	Product        Product `json:"product" gorm:"foreignKey:ProductID"`
	Quantity       int     `json:"quantity" gorm:"not null"`
	PriceWhenAdded float64 `json:"price_when_added" gorm:"not null"`
}

type CartCreate struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type CartUpdate struct {
	Quantity int `json:"quantity" binding:"gte=0"`
}
