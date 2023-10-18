package domain

type Coupons struct {
	Id                 int     `json:"coupon_id" gorm:"uniquekey; not null"`
	Coupon             string  `json:"coupon" gorm:"coupon"`
	DiscountPercentage int     `json:"discount_percentage"`
	Validity           bool    `json:"validity"`
	MinimumPrice       float64 `json:"minimum_price"`
}
