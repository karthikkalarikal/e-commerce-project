package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type OrderRepositry interface {
	AddToOrder(int, int) (domain.Order, error)
	GetOrder(int) (domain.Order, error)
	GetDeliveryAddress(userId int) (int, error)
	GetUserOrders(userId int) ([]models.Cart, error)
	ChangeOrderStatus(orderId int) (domain.Order, error)
	TotalAmountInCart(userId int) (float64, error)
	AddAmountToOrder(amount float64, orderId uint) error
	GetDetailedOrderThroughId(orderId int) (models.CombinedOrderDetails, error)
	GetPaymentStatus(orderId int) (bool, error)
	GetTotalAmount(orderId int) (domain.Order, error)
	AddMoneyToWallet(userId int, amount float64) (domain.Wallet, error)
	CheckForWallet(userId int) (bool, error)
	AddMondyToExistingWallet(userId int, amount float64) (domain.Wallet, error)
	GetWalletByUserId(userId int) (domain.Wallet, error)
	GetItemsByOrderId(orderId int) ([]models.ItemDetails, error)
}
