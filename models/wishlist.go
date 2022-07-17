package models

type Wishlist struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserID    string  `json:"-"`
	User      Profile `json:"-" gorm:"foreignKey:UserID"`
	ProductID uint    `json:"-"`
	Product   Product `json:"product" gorm:"foreignKey:ProductID"`
}

type WishlistCreate struct {
	ProductID uint `json:"product_id" binding:"required"`
}
