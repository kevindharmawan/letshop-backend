package models

type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"not null"`
	Desc string `json:"description" gorm:"not null"`
}
