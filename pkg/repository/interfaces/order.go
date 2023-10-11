package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/domain"

type OrderRepositry interface {
	AddToOrder(int, int) error
	GetOrder(int) (domain.Order, error)
}
