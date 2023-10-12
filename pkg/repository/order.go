package repository

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
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

func (repo *orderRepositryImpl) AddToOrder(addressId, cartId int) error {
	query := `
		insert into orders(user_id,address_id,cart_id)
		select c.user_id,c.id, a.address_id
		from carts c
		join addresses a on c.user_id = a.user_id
		where a.address_id = $1 and c.id = $2

	`
	if err := repo.db.Exec(query, addressId, cartId).Error; err != nil {
		return err
	}
	return nil
}

// -------------------------------------------- get order table by user id ----------------------------------\\

func (repo *orderRepositryImpl) GetOrder(userId int) (domain.Order, error) {
	var body domain.Order
	query := `
		select * from orders
		where user_id = $1
	`
	if err := repo.db.Raw(query, userId).Scan(&body).Error; err != nil {
		return domain.Order{}, err
	}

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

// func (repo *orderRepositryImpl) GetCartsOfUser (userId int) ([]int, error)
