package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type CartRepository interface {
	AddToCart(cart models.CartItems, cartId int) (models.CartItems, error)
	CartItemListing(int) ([]models.CartItems, error)
	CartItemQuantityUpdations(int, int, string) error
	CartItemDeletion(int, int) error
	MakeNewCart(int) (models.Cart, error)
}
