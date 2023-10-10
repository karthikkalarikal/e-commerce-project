package usecase

import (
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

func (usecase *cartUseCaseImpl) AddToCart(cart domain.Cart, userId int, productId int) (domain.Cart, error) {

	// to do check for stocks

	body, err := usecase.repo.AddToCart(cart, userId, productId)
	if err != nil {
		return domain.Cart{}, err
	}

	return body, nil
}

// -------------------------------------------- cart item listing ----------------------------------------------- \\

func (usecase *cartUseCaseImpl) CartItemListing(userId int) ([]models.CartItems, error) {
	body, err := usecase.repo.CartItemListing(userId)
	if err != nil {
		return []models.CartItems{}, err
	}

	return body, nil
}

// ----------------------------------------------cart item quantity updation-----------------------------------------\\

func (usecase *cartUseCaseImpl) CartItemQuantityUpdations(userId, productInt int, quantity string) ([]models.CartItems, error) {
	if err := usecase.repo.CartItemQuantityUpdations(userId, productInt, quantity); err != nil {
		return []models.CartItems{}, err
	}

	body, err := usecase.repo.CartItemListing(userId)
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

	body, err := usecase.repo.CartItemListing(userId)
	if err != nil {
		return []models.CartItems{}, err
	}

	return body, nil
}
