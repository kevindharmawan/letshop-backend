package models

type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Desc string `json:"description" gorm:"not null"`
}

type CategoryCreate struct {
	Name string `json:"name" binding:"required"`
	Desc string `json:"description" binding:"required"`
}

type CategoryUpdate struct {
	Name string `json:"name" binding:"-"`
	Desc string `json:"description" binding:"-"`
}
