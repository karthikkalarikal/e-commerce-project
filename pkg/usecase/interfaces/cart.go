package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type CartUseCase interface {
	AddToCart(domain.Cart, int, int) (domain.Cart, error)
	CartItemListing(int) ([]models.CartItems, error)
	CartItemQuantityUpdations(int, int, string) ([]models.CartItems, error)
	CartItemDeletion(int, int) ([]models.CartItems, error)
}
