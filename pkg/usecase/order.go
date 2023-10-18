package usecase

import (
	"fmt"

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
	fmt.Println("addressid", addressId)

	totlaAmount, err := repo.orderRepo.TotalAmountInCart(cartId)
	if err != nil {
		return domain.Order{}, nil
	}
	fmt.Println("total amount", totlaAmount)

	body, err := repo.orderRepo.AddToOrder(cartId, addressId)
	if err != nil {
		return domain.Order{}, err
	}

	fmt.Println("body", body)
	if err := repo.orderRepo.AddAmountToOrder(totlaAmount, body.ID); err != nil {
		return domain.Order{}, err
	}
	body2, err := repo.orderRepo.GetOrder(int(body.ID))
	if err != nil {
		return domain.Order{}, err
	}
	fmt.Println("body2", body2)

	return body2, nil

}

// --------------------------------------------- cancel orders -------------------------------------- \\

func (repo *orderUseCaseImpl) CancelOrder(orderId int) (domain.Order, domain.Wallet, error) { // change the domain order to model cancelled order after
	payment, err := repo.orderRepo.GetPaymentStatus(orderId) // get payment status
	if err != nil {
		return domain.Order{}, domain.Wallet{}, nil
	}

	if payment { // if payed the amount should be debited to wallet
		bodyOrder, err := repo.orderRepo.GetTotalAmount(orderId) // get body of order by order id
		if err != nil {
			return domain.Order{}, domain.Wallet{}, err
		}
		okWallet, err := repo.orderRepo.CheckForWallet(bodyOrder.UserId) // if the user has a wallet already
		if err != nil {
			return domain.Order{}, domain.Wallet{}, nil
		}
		if !okWallet {
			bodyWallet, err := repo.orderRepo.AddMoneyToWallet(bodyOrder.UserId, bodyOrder.Amount) // in the case of no wallet
			if err != nil {
				return domain.Order{}, domain.Wallet{}, err
			}
			orderRetBody, err := repo.orderRepo.ChangeOrderStatus(orderId) // change the order and payment status
			if err != nil {
				return domain.Order{}, domain.Wallet{}, err
			}

			return orderRetBody, bodyWallet, nil
		} else {
			bodyWallet, err := repo.orderRepo.AddMondyToExistingWallet(bodyOrder.UserId, bodyOrder.Amount) // add money to existing wallet
			if err != nil {
				return domain.Order{}, domain.Wallet{}, err
			}
			orderRetBody, err := repo.orderRepo.ChangeOrderStatus(orderId) // change order status again
			if err != nil {
				return domain.Order{}, domain.Wallet{}, err
			}

			return orderRetBody, bodyWallet, nil
		}

	}
	orderRetBody, err := repo.orderRepo.ChangeOrderStatus(orderId) // in the case that the order is not payed yet
	if err != nil {
		return domain.Order{}, domain.Wallet{}, err
	}
	walletBody, err := repo.orderRepo.GetWalletByUserId(orderRetBody.UserId) // get wallet
	if err != nil {
		return domain.Order{}, domain.Wallet{}, err
	}

	return orderRetBody, walletBody, nil
}

// ------------------------------------------ get orders of user ----------------------------------- \\

func (repo *orderUseCaseImpl) ViewOrder(orderId int) (models.CombinedOrderDetails, error) {

	body, err := repo.orderRepo.GetDetailedOrderThroughId(orderId)
	if err != nil {
		return models.CombinedOrderDetails{}, err
	}
	return body, err
}

// ------------------------------------ display the wallet of user by demand ------------------------------ \\ 

func (repo *orderUseCaseImpl) ViewWalletByUserId(userId int) (domain.Wallet, error) {
	body, err := repo.orderRepo.GetWalletByUserId(userId)
	if err != nil {
		return domain.Wallet{}, err
	}
	return body, nil
}
