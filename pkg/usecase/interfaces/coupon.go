package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type CouponUseCase interface {
	AddCoupon(coupon string, discount int, validity bool, minprice float64) (domain.Coupons, error)
	ViewCoupon() ([]models.CouponInput, error)
	ExpireCoupon(name string) (models.CouponInput, error)
	RedeemCoupon(string, int) (models.CombinedOrderDetails, error)
}
