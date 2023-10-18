package models

type CouponInput struct {
	Coupon             string  `json:"coupon" gorm:"not null"`
	DiscountPercentage int     `json:"discount_percentage" gorm:"not null"`
	Validity           bool    `json:"validity"`
	MinimumPrice       float64 `json:"minimum_price" gorm:"not null"`
}
