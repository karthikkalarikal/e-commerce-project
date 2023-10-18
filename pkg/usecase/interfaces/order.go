package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type OrderUseCase interface {
	AddToOrder(userId, cartId int) (domain.Order, error)
	ViewOrder(orderId int) (models.CombinedOrderDetails, error)
}
