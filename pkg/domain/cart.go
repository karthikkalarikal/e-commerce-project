package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model `json:"-"`
	UserId     int     `json:"user_id" gorm:"not null"`
	Users      Users   `json:"-" gorm:"foreignkey:UserId"`
	ProductId  int     `json:"product_id"`
	Product    Product `json:"-" gorm:"foreignkey:ProductId"`
	Quantity   float64 `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
