package usecase

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	repo "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
)

type OrderUseCaseImpl struct {
	orderRepo repo.OrderRepositry
}

func NewOrderUseCase(repo repo.OrderRepositry) interfaces.OrderUseCase {
	return &OrderUseCaseImpl{
		orderRepo: repo,
	}
}

// --------------------------------------------- add to order ----------------------------------\\

func (repo *OrderUseCaseImpl) AddToOrder(userId, cartId int) (domain.Order, error) {

	addressId, err := repo.orderRepo.GetDeliveryAddress(userId)
	if err != nil {
		return domain.Order{}, nil
	}

	if err := repo.orderRepo.AddToOrder(cartId, addressId); err != nil {
		return domain.Order{}, err
	}
	body, err := repo.orderRepo.GetOrder(userId)
	if err != nil {
		return domain.Order{}, err
	}

	return body, nil

}
