package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	// CartId int   `json:"cart_id" gorm:"primarykey;not null"`
	UserId int   `json:"user_id" gorm:"not null"`
	Users  Users `json:"-" gorm:"foreignkey:UserId"`
}

type CartItems struct {
	CartItemsId int     `json:"cartitems_id" gorm:"primarykey;not null"`
	CartId      int     `json:"cart_id" gorm:"not null"`
	Cart        Cart    `json:"-" gorm:"foreignkey:CartId"`
	ProductId   int     `json:"product_id"`
	Product     Product `json:"-" gorm:"foreignkey:ProductId"`
	Quantity    float64 `json:"quantity"`
	Amount      float64 `json:"amount"`
}
