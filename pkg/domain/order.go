package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId      int    `json:"user_id" gorm:"not null"`
	AddressID   int    `json:"address_id" gorm:"not null"`
	CartId      int    `json:"cart_id" gorm:"not null"`
	OrderStatus string `json:"order_status" gorm:"default:'confirmed'"`
}

type Wallet struct {
	WalletId int   `json:"wallet_id" gorm:"not null"`
	UserId   int   `json:"user_id" gorm:"not null"`
	Users    Users `json:"-" gorm:"foreignkey:UserId"`
	CartId   int   `json:"cart_id" gorm:"not null"`
	Cart     Cart  `json:"-" gorm:"foreignkey:CartId"`
}
