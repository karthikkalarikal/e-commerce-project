package usecase

import (
	"fmt"

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

func (usecase *cartUseCaseImpl) CartItemQuantityUpdations(userId, cartInt int, quantity string) ([]models.CartItems, error) {
	if err := usecase.repo.CartItemQuantityUpdations(userId, cartInt, quantity); err != nil {
		return []models.CartItems{}, err
	}

	body, err := usecase.repo.CartItemListing(userId, cartInt)
	if err != nil {
		return []models.CartItems{}, err
	}

	return body, nil

}

// ----------------------------------------------cart item deletion ----------------------------------------------------\\
func (usecase *cartUseCaseImpl) CartItemDeletion(userId, productInt int) ([]models.CartItems, error) {
	if err := usecase.repo.CartItemDeletion(userId, productInt); err != nil {
		return []models.CartItems{}, err
	}

	body, err := usecase.repo.CartItemListing(userId, productInt)
	if err != nil {
		return []models.CartItems{}, err
	}

	return body, nil
}
