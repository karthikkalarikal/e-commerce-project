package usecase

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type couponUseCaseImpl struct {
	repo interfaces.CouponRepository
}

func NewCouponUseCase(repo interfaces.CouponRepository) usecase.CouponUseCase {
	return &couponUseCaseImpl{
		repo: repo,
	}
}

// -------------------------------------------------- add coupons ---------------------------------------------- \\

func (repo *couponUseCaseImpl) AddCoupon(coupon string, discount int, validity bool, minprice float64) (domain.Coupons, error) {

	// validation goes here

	// insert coupon

	body, err := repo.repo.AddCoupon(coupon, discount, validity, minprice)
	if err != nil {
		return domain.Coupons{}, err
	}

	return body, nil
}

// ------------------------------------------------------- view coupons --------------------------------------------- \\

func (repo *couponUseCaseImpl) ViewCoupon() ([]models.CouponInput, error) {
	body, err := repo.repo.ViewCoupon()

	if err != nil {
		return []models.CouponInput{}, err
	}
	return body, nil

}
