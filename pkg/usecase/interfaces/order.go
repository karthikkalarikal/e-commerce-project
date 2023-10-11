package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/domain"

type OrderUseCase interface {
	AddToOrder(userId, cartId, addressId int) (domain.Order, error)
}
