package usecase

import (
	"errors"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	repo "github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type orderUseCaseImpl struct {
	orderRepo repo.OrderRepositry
	helpRepo  repo.HelperRepository
}

func NewOrderUseCase(repo repo.OrderRepositry, helprepo repo.HelperRepository) interfaces.OrderUseCase {
	return &orderUseCaseImpl{
		orderRepo: repo,
		helpRepo:  helprepo,
	}
}

// --------------------------------------------- add to order ----------------------------------\\

func (repo *orderUseCaseImpl) AddToOrder(userId, cartId int) (domain.Order, error) {

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

// --------------------------------------------- cancel orders -------------------------------------- \\

func (repo *orderUseCaseImpl) CancelOrder(userId int) ([]models.Cart, error) {
	ok, err := repo.helpRepo.FindIfUserExists(userId, "orders") // validating user id
	if err != nil {
		return []models.Cart{}, err
	}
	if !ok {
		return []models.Cart{}, errors.New("no user with this id")
	}

	body, err := repo.orderRepo.GetUserOrders(userId) // getting all the orders of user
	if err != nil {
		return []models.Cart{}, err
	}
	for _, v := range body {
		if err := repo.orderRepo.ChangeStatus(v.CartId); err != nil {
			return []models.Cart{}, err
		}
	}

	return body, nil

}
