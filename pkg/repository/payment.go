package repository

import (
	"errors"
	"fmt"

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

// --------------------------------------- add payment details ----------------------------------------- \\

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

// ---------------------------------------- update payment details ------------------------------------------- \\

func (repo *paymentRepositoryImpl) UpdatePaymentDetails(orderId string, paymentId string) error {
	fmt.Println("razerId,paymetnId", orderId, paymentId)
	if err := repo.DB.Exec("update razer_pays set payment = $1 where razer_id = $2", paymentId, orderId).Error; err != nil {
		err = errors.New("error in updating the razer pay table " + err.Error())
		return err
	}
	return nil
}

// ------------------------------------------- check payment status ----------------------------------- \\

func (repo *paymentRepositoryImpl) GetPaymentStatus(orderId string) (bool, error) {
	var body bool
	err := repo.DB.Raw("select payment_status from orders where id = $1", orderId).Scan(&body).Error
	if err != nil {
		return false, err
	}
	return body, nil
}

// -------------------------------------------- update payment status ---------------------------------- \\

func (repo *paymentRepositoryImpl) UpdatePaymentStatus(status bool, orderId string) error {

	query := `
			update orders set payment_status = $1 where id = $2 
	`
	if err := repo.DB.Exec(query, status, orderId).Error; err != nil {
		err = errors.New("error in updating orders payment status " + err.Error())
		return err
	}
	return nil
}
