package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type OrderRepositry interface {
	AddToOrder(int, int) (domain.Order, error)
	GetOrder(int) (domain.Order, error)
	GetDeliveryAddress(userId int) (int, error)
	GetUserOrders(userId int) ([]models.Cart, error)
	ChangeStatus(userId int) error
	TotalAmountInCart(userId int) (float64, error)
	AddAmountToOrder(amount float64, orderId uint) error
}
