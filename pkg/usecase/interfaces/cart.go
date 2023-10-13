package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type CartUseCase interface {
	AddToCart(cartitems models.CartItems, userId, cartId int) (models.CartItems, error)
	CartItemListing(int, int) ([]models.CartItems, error)
	CartItemQuantityUpdations(int, int) (models.CartItems, error)
	CartItemDeletion(int) (models.CartItems, error)
}
