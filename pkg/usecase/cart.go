package usecase

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	repository "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type cartUseCaseImpl struct {
	repo repository.CartRepository
}

func NewCartUseCase(usecase repository.CartRepository) interfaces.CartUseCase {
	return &cartUseCaseImpl{
		repo: usecase,
	}
}

// ------------------------------------------add to cart usecase ---------------------------------------------- \\

func (usecase *cartUseCaseImpl) AddToCart(cartitems models.CartItems, userId, cartId int) (models.CartItems, error) {

	if cartId <= 0 { // incase the cart isnt specified
		newCart, err := usecase.repo.MakeNewCart(userId)
		if err != nil {
			return models.CartItems{}, err
		}
		cartId = newCart.CartId
		fmt.Println("cart id", cartId, newCart)
	}

	// to do check for stocks

	body, err := usecase.repo.AddToCart(cartitems, cartId) // adds the items into cart items database as specified
	if err != nil {
		return models.CartItems{}, err
	}

	return body, nil
}

// -------------------------------------------- cart item listing ----------------------------------------------- \\

func (usecase *cartUseCaseImpl) CartItemListing(userId, cartId int) ([]models.CartItems, error) {
	body, err := usecase.repo.CartItemListing(userId, cartId)
	if err != nil {
		return []models.CartItems{}, err
	}

	return body, nil
}

// ----------------------------------------------cart item quantity updation-----------------------------------------\\

func (usecase *cartUseCaseImpl) CartItemQuantityUpdations(cartItems, quantity int) (models.CartItems, error) {
	if err := usecase.repo.CartItemQuantityUpdations(cartItems, quantity); err != nil {
		return models.CartItems{}, err
	}

	body, err := usecase.repo.CartItemsById(cartItems)
	if err != nil {
		return models.CartItems{}, err
	}

	return body, nil

}

// ----------------------------------------------cart item deletion ----------------------------------------------------\\
func (usecase *cartUseCaseImpl) CartItemDeletion(cartItemId int) (models.CartItems, error) {
	body, err := usecase.repo.CartItemsById(cartItemId)
	if err != nil {
		return models.CartItems{}, err
	}

	err = usecase.repo.CartItemDeletion(cartItemId)
	if err != nil {
		return models.CartItems{}, err
	}

	return body, nil
}

// ---------------------------------------------- check if user has carts ---------------------------------------- \\

func (usecase *cartUseCaseImpl) CheckUserCartById(userInt int) error {

	err := usecase.repo.CheckUserCartById(userInt)
	if err != nil {
		return err
	}

	return nil

}

// -------------------------------------------- carts of users ------------------------------------------- \\

func (usecase *cartUseCaseImpl) GetCartsByUserId(userInt int) ([]domain.Cart, error) {

	body, err := usecase.repo.GetCartsByUserId(userInt)
	if err != nil {
		return []domain.Cart{}, err
	}
	return body, nil
}
