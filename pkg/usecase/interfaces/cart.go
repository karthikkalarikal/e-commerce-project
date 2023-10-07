package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/domain"

type CartUseCase interface {
	AddToCart(domain.Cart, int) (domain.Cart, error)
}
