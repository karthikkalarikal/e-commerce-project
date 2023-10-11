package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId      int    `json:"user_id" gorm:"not null"`
	AddressID   int    `json:"address_id" gorm:"not null"`
	CartId      int    `json:"cart_id" gorm:"not null"`
	OrderStatus string `json:"order_status" gorm:"default:'pending'"`
}
