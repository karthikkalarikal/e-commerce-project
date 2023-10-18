package usecase

import (
	"errors"
	"fmt"

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

// --------------------------------------------------- expire coupons by name --------------------------------------- \\

func (repo *couponUseCaseImpl) ExpireCoupon(name string) (models.CouponInput, error) {
	exists, err := repo.repo.CheckCoupon(name) // check if it exists

	if err != nil {

		return models.CouponInput{}, err
	}

	if exists {

		isValid, err := repo.repo.CheckCouponValidity(name) // check if it is valid or not

		if err != nil {

			return models.CouponInput{}, err

		}

		if isValid {
			fmt.Println(isValid)
			body, err := repo.repo.ExpireCoupon(name) // modify the value of it validity

			if err != nil {

				err = errors.New("error expiring the token" + err.Error())
				return models.CouponInput{}, err

			}
			fmt.Println(body)
			return body, nil

		}

		return models.CouponInput{}, errors.New("the coupon is already expired")

	}

	return models.CouponInput{}, errors.New("coupon is not available")
}
