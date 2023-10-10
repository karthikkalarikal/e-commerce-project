package usecase

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	repository "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
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
