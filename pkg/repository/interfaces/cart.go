package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type CartRepository interface {
	AddToCart(cart models.CartItems, cartId int) (models.CartItems, error)
	CartItemListing(userId, cartId int) ([]models.CartItems, error)
	CartItemQuantityUpdations(int, int) error
	CartItemDeletion(cartItemsId int) error
	MakeNewCart(int) (models.Cart, error)
	CartItemsById(cartItemsId int) (models.CartItems, error)
	CheckUserCartById(userId int) error
	GetCartsByUserId(userInt int) ([]domain.Cart, error)
}
