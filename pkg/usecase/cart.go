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

// add to cart usecase
func (usecase *cartUseCaseImpl) AddToCart(cart domain.Cart, id int) (domain.Cart, error) {

	body, err := usecase.repo.AddToCart(cart, id)
	if err != nil {
		return domain.Cart{}, err
	}

	return body, nil
}
