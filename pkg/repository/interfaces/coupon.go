package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type CouponRepository interface {
	AddCoupon(coupon string, discount int, validity bool, minprice float64) (domain.Coupons, error)
	ViewCoupon() ([]models.CouponInput, error)
	ExpireCoupon(coupon string) (models.CouponInput, error)
	CheckCouponValidity(coupon string) (bool, error)
	CheckCoupon(coupon string) (bool, error)
}
