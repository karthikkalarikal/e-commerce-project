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
	repo      interfaces.CouponRepository
	orderRepo interfaces.OrderRepositry
}

func NewCouponUseCase(repo interfaces.CouponRepository, order interfaces.OrderRepositry) usecase.CouponUseCase {
	return &couponUseCaseImpl{
		repo:      repo,
		orderRepo: order,
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

// -------------------------------------- redeem coupons --------------------------------------- \\

func (repo *couponUseCaseImpl) RedeemCoupon(coupon string, orderId int) (models.CombinedOrderDetails, error) {

	couponDetails, err := repo.repo.GetCoupon(coupon) // get coupon details
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}
	orderDetails, err := repo.orderRepo.GetOrder(orderId) // get order details
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}
	if couponDetails.Validity && couponDetails.MinimumPrice < orderDetails.Amount {

		amount := orderDetails.Amount - (orderDetails.Amount * float64(couponDetails.DiscountPercentage) / 100)
		fmt.Println("amount changed ", amount)
		err := repo.repo.ChangeOrderAmount(orderId, amount)
		if err != nil {
			return models.CombinedOrderDetails{}, err
		}
		_, err = repo.repo.ExpireCoupon(coupon)
		if err != nil {
			return models.CombinedOrderDetails{}, err
		}
		body, err := repo.orderRepo.GetDetailedOrderThroughId(orderId)
		if err != nil {
			return models.CombinedOrderDetails{}, err
		}
		return body, nil
	}
	return models.CombinedOrderDetails{}, errors.New("the coupon is not valid")
}
