package repository

import (
	"errors"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type couponRepositoryImpl struct {
	DB *gorm.DB
}

func NewCouponRepository(db *gorm.DB) interfaces.CouponRepository {
	return &couponRepositoryImpl{
		DB: db,
	}
}

// ------------------------------------add coupon by admin ------------------------------------ \\

func (repo *couponRepositoryImpl) AddCoupon(coupon string, discount int, validity bool, minprice float64) (domain.Coupons, error) {

	var body domain.Coupons

	query := `
	insert into coupons(coupon,discount_percentage,validity,minimum_price) 
	values($1,$2,$3,$4) returning*
	
	` // the query to insert values into coupons table
	if err := repo.DB.Raw(query, coupon, discount, validity, minprice).Scan(&body).Error; err != nil { // its recomended to use exec but raw works here and you could get the values back only in raw
		err = errors.New("error in inserting into coupons table" + err.Error())
		return domain.Coupons{}, err
	}

	return body, nil
}

// ------------------------------------ view coupon by admin ---------------------------------------- \\

func (repo *couponRepositoryImpl) ViewCoupon() ([]models.CouponInput, error) {
	var body []models.CouponInput

	query := `
		select * from coupons
	
	`
	if err := repo.DB.Raw(query).Scan(&body).Error; err != nil {
		return []models.CouponInput{}, err
	}
	return body, nil
}

// -------------------------------------- expire coupon ------------------------------------- \\

func (repo *couponRepositoryImpl) ExpireCoupon(coupon string) (models.CouponInput, error) {
	var body models.CouponInput

	query := `
	update coupons set validity = false where coupon = $1 returning *
	`
	if err := repo.DB.Raw(query, coupon).Scan(&body).Error; err != nil {
		return models.CouponInput{}, err
	}
	return body, nil
}

// -------------------------------------- check coupon for expiry ------------------------------ \\

func (repo *couponRepositoryImpl) CheckCouponValidity(coupon string) (bool, error) {
	var body bool

	query := `
		select validity from coupons where coupon = $1 
	`
	if err := repo.DB.Raw(query, coupon).Scan(&body).Error; err != nil {
		return false, err
	}
	return body, nil
}

// ---------------------------------------- check coupon by name ------------------------------------ \\

func (repo *couponRepositoryImpl) CheckCoupon(coupon string) (bool, error) {
	body := 0
	query := `
	select count(*) from coupons where coupon = $1
	
	`
	if err := repo.DB.Raw(query, coupon).Scan(&body).Error; err != nil {
		return false, err
	}
	return body > 0, nil
}
