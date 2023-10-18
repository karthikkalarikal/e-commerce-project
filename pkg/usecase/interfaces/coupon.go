package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/domain"

type CouponUseCase interface {
	AddCoupon(coupon string, discount int, validity bool, minprice float64) (domain.Coupons, error)
}
