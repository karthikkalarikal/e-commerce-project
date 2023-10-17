package repository

import (
	"errors"

	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepositoryImpl{
		DB: DB,
	}
}

func (repo *paymentRepositoryImpl) AddRazorPayDetails(orderId int, razorPayId string) error {
	query := `
	insert into razer_pays (order_id,razer_id) values($1,$2) 
	`
	if err := repo.DB.Exec(query, orderId, razorPayId).Error; err != nil {
		err = errors.New("error in inserting values to razor pay data table" + err.Error())
		return err
	}
	return nil
}
