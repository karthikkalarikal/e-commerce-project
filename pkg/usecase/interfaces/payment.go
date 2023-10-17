package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type PaymentUseCase interface {
	MakePaymentRazorpay(orderId, userId int) (models.CombinedOrderDetails, string, error)
}
