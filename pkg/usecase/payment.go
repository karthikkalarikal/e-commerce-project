package usecase

import (
	"errors"
	"fmt"

	payment "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"github.com/razorpay/razorpay-go"
)

type paymentUsecaseImpl struct {
	paymentRepo     payment.PaymentRepository
	orderRepository usecase.OrderRepositry
}

func NewPaymentUseCase(repo usecase.OrderRepositry, payment payment.PaymentRepository) interfaces.PaymentUseCase {
	return &paymentUsecaseImpl{
		orderRepository: repo,
		paymentRepo:     payment,
	}
}

// ---------------------------------------- make payment through razor pay --------------------------------------- \\

func (repo *paymentUsecaseImpl) MakePaymentRazorpay(orderId, userId int) (models.CombinedOrderDetails, string, error) {

	order, err := repo.orderRepository.GetOrder(orderId)
	if err != nil {
		err = errors.New("error in getting order details through order id" + err.Error())
		return models.CombinedOrderDetails{}, "", err
	}

	client := razorpay.NewClient("rzp_test_YSrXzAXNevri34", "zH7CC1BlsmjsVoWGZBXr8FeL")

	fmt.Println("order amount", order.Amount)
	data := map[string]interface{}{
		"amount":   int(order.Amount) * 100,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}

	body, err := client.Order.Create(data, nil)
	if err != nil {
		return models.CombinedOrderDetails{}, "", nil
	}
	fmt.Println("body", body)
	razorPayOrderId := body["id"].(string)

	err = repo.paymentRepo.AddRazorPayDetails(orderId, razorPayOrderId)
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}
	body2, err := repo.orderRepository.GetDetailedOrderThroughId(int(order.ID))
	if err != nil {
		return models.CombinedOrderDetails{}, "", err
	}
	fmt.Println("body 2 usecase", body2)

	return body2, razorPayOrderId, nil
}

// ------------------------------------------------- verify payment razor pay ------------------------------------ \\

func (repo *paymentUsecaseImpl) SavePaymentDetails(paymentId, razorId, orderId string) error {

	status, err := repo.paymentRepo.GetPaymentStatus(orderId)
	if err != nil {
		return err
	}
	fmt.Println("status", status)
	if !status {
		err = repo.paymentRepo.UpdatePaymentDetails(razorId, paymentId)
		if err != nil {
			return err
		}

		err = repo.paymentRepo.UpdatePaymentStatus(true, orderId)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("already paid")

}
