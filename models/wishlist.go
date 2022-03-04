package models

type Wishlist struct {
	UserID  string  `json:"user_id" gorm:"type:varchar(30);primaryKey"`
	Product Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;primaryKey"`
}
