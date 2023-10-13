package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type OrderRepositry interface {
	AddToOrder(int, int) error
	GetOrder(int) (domain.Order, error)
	GetDeliveryAddress(userId int) (int, error)
	GetUserOrders(userId int) ([]models.Cart, error)
	ChangeStatus(userId int) error
}
