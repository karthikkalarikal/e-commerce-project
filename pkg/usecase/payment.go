package usecase

import (
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
)

type paymentUsecaseImpl struct {
	orderRepository usecase.OrderRepositry
}

func NewPaymentUseCase(repo usecase.OrderRepositry) interfaces.PaymentUseCase {
	return &paymentUsecaseImpl{
		orderRepository: repo,
	}
}

// func (repo *paymentUsecaseImpl) MakePaymentRazorpay() (models.CombinedOrderDetails, error) {
// 	client := razorpay.NewClient("rzp_test_YSrXzAXNevri34", "zH7CC1BlsmjsVoWGZBXr8FeL")

// 	data := map[string]interface{}{
// 		"amount":   int() * 100,
// 		"currency": "INR",
// 		"receipt":  "some_receipt_id",
// 	}

// 	body, err := client.Order.Create(data, nil)
// 	if err != nil {
// 		return models.CombinedOrderDetails{}, nil
// 	}

// 	return

// }
