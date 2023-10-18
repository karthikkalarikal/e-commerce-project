package repository

import (
	"errors"
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type orderRepositryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepositry {
	return &orderRepositryImpl{
		db: db,
	}
}

// ------------------------------------------- add to order table -------------------------------------\\

func (repo *orderRepositryImpl) AddToOrder(addressId, cartId int) (domain.Order, error) {

	var body domain.Order
	query := `
		insert into orders(user_id,address_id,cart_id)
		select c.user_id,c.id, a.address_id
		from carts c
		join addresses a on c.user_id = a.user_id
		where a.address_id = $1 and c.id = $2 
		returning *
	`
	if err := repo.db.Raw(query, addressId, cartId).Scan(&body).Error; err != nil {
		err = errors.New("failed to add order to the data base" + err.Error())
		return domain.Order{}, err
	}
	return body, nil
}

// ------------------------------------------- add amount to order table ---------------------------------- \\

func (repo *orderRepositryImpl) AddAmountToOrder(amount float64, orderId uint) error {
	query := `
		update orders set amount = $1 where id = $2
	`

	if err := repo.db.Exec(query, amount, orderId).Error; err != nil {
		return err
	}
	return nil
}

// -------------------------------------------- get order table by order id ----------------------------------\\

func (repo *orderRepositryImpl) GetOrder(orderId int) (domain.Order, error) {
	var body domain.Order
	query := `
		select * from orders
		where id = $1
	`
	if err := repo.db.Raw(query, orderId).Scan(&body).Error; err != nil {
		return domain.Order{}, err
	}
	fmt.Println("amount", body.Amount)
	return body, nil
}

// ------------------------------------------ get selected address from user id ------------------------------------- \\

func (repo *orderRepositryImpl) GetDeliveryAddress(userId int) (int, error) {
	var address domain.Address

	query := `
		select address_id from addresses
		where selection = true and user_id = $1
	`
	if err := repo.db.Raw(query, userId).Scan(&address).Error; err != nil {
		return 0, err
	}
	fmt.Println("addresses", address.AddressId)
	return address.AddressId, nil
}

// ------------------------------------- get carts of user from user id -------------------------------------------------- \\

func (repo *orderRepositryImpl) GetUserOrders(userId int) ([]models.Cart, error) {
	var body []models.Cart

	query := `
	select * from carts where user_id = $1 and status = 'confirmed'
	
	` // get all the carts that have the given id and the status of confirmed
	if err := repo.db.Raw(query, userId).Scan(&body).Error; err != nil {
		return []models.Cart{}, err
	}

	return body, nil
}

// --------------------------------------- change order status into cancel ------------------------------------------------ \\

func (repo *orderRepositryImpl) ChangeOrderStatus(orderId int) (domain.Order, error) {
	var body domain.Order
	query := `
	update orders set order_status = 'cancel' , payment_status = false where id = $1 returning *
`
	if err := repo.db.Raw(query, orderId).Scan(&body).Error; err != nil {
		return domain.Order{}, err
	}
	return body, nil
}

// ----------------------------------------- get total amount of a cart --------------------------------------------------- \\

func (repo *orderRepositryImpl) TotalAmountInCart(cartId int) (float64, error) {
	var amounts []float64

	query := `
		select ci.amount from carts c 
		join cart_items ci
		on c.id = ci.cart_id 
		where c.cart_id = $1
	`
	if err := repo.db.Raw(query, cartId).Pluck("amount", &amounts).Error; err != nil {
		return 0, err
	}
	sum := 0.0
	for _, v := range amounts {

		sum += v
	}
	return sum, nil
}

// ---------------------------------------- get full order details through order id --------------------------------------------- \\

func (repo *orderRepositryImpl) GetDetailedOrderThroughId(orderId int) (models.CombinedOrderDetails, error) {
	var body models.CombinedOrderDetails

	query := `
	select 
		o.id as order_id,
		o.amount as amount,
		o.order_status as order_status,
		o.payment_status as payment_status,
		u.name as name,
		u.email as email,
		u.phone as phone,
		a.house_name as house_name,
		a.state as state,
		a.pin as pin,
		a.street as street,
		a.city as city
	from orders o
	join users u on o.user_id = u.user_id
	join addresses a on o.address_id = a.address_id 
	where o.id = $1
	`
	if err := repo.db.Raw(query, orderId).Scan(&body).Error; err != nil {
		err = errors.New("error in getting detailed order through id in repository" + err.Error())
		return models.CombinedOrderDetails{}, err
	}
	fmt.Println("body in repo", body.Amount)
	return body, nil
}

// --------------------------------- check payment status by order id --------------------------------- \\

func (repo *orderRepositryImpl) GetPaymentStatus(orderId int) (bool, error) {
	var body bool
	query := `select payment_status from orders where id = $1`

	if err := repo.db.Raw(query, orderId).Scan(&body).Error; err != nil {
		return false, err
	}
	return body, nil
}

// ----------------------------------- get total amount ------------------------------------------ \\

func (repo *orderRepositryImpl) GetTotalAmount(orderId int) (domain.Order, error) {
	var amount domain.Order

	query := `
	select * from orders where id = $1
	`
	if err := repo.db.Raw(query, orderId).Scan(&amount).Error; err != nil {
		err = errors.New("error in getting total amount")
		return domain.Order{}, err
	}
	return amount, nil

}

// -------------------------------------- check if user has a wallet ------------------------------- \\

func (repo *orderRepositryImpl) CheckForWallet(userId int) (bool, error) {
	var body int
	query := `select count(*) from wallets where user_id = $1`

	if err := repo.db.Raw(query, userId).Scan(&body).Error; err != nil {
		return false, err
	}
	return body > 0, nil
}

// -------------------------------------- add money to wallet ----------------------------------------- \\

func (repo *orderRepositryImpl) AddMoneyToWallet(userId int, amount float64) (domain.Wallet, error) {
	var body domain.Wallet

	query := `
	insert into wallets(user_id,amount) values($1,$2) returning *
	`
	if err := repo.db.Raw(query, userId, amount).Scan(&body).Error; err != nil {
		err = errors.New("error in inserting into wallets table")
		return domain.Wallet{}, err
	}

	return body, nil
}

// ----------------------------------------- add money to existing wallet ----------------------------------- \\

func (repo *orderRepositryImpl) AddMondyToExistingWallet(userId int, amount float64) (domain.Wallet, error) {
	var body domain.Wallet
	var balance float64

	query := `
		select amount from wallets where user_id = $1
	`
	if err := repo.db.Raw(query, userId).Scan(&balance).Error; err != nil {
		err = errors.New("error in getting balance from wallet by user id")
		return domain.Wallet{}, err
	}
	balance = balance + amount
	query2 := `
	update wallets set amount = $1 where user_id = $2 returning *
	`
	if err := repo.db.Raw(query2, balance, userId).Scan(&body).Error; err != nil {
		err = errors.New("error in updating wallet value")
		return domain.Wallet{}, err
	}
	return body, nil
}

func (repo *orderRepositryImpl) GetWalletByUserId(userId int) (domain.Wallet, error) {
	var body domain.Wallet
	query := `
		select * from wallets where user_id = $1
	
	`
	if err := repo.db.Raw(query, userId).Scan(&body).Error; err != nil {
		return domain.Wallet{}, err
	}
	return body, nil
}
