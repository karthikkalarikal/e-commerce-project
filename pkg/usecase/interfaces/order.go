package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/domain"

type OrderUseCase interface {
	AddToOrder(userId, cartId int) (domain.Order, error)
}
