package repository

import (
	"errors"
	"fmt"
	"log"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type cartRepositoryImpl struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) interfaces.CartRepository {
	return &cartRepositoryImpl{
		DB: db,
	}
}

// ------------------------------------------ make a newcart ---------------------------------------------------- \\

func (repo *cartRepositoryImpl) MakeNewCart(userId int) (models.Cart, error) {
	var cart models.Cart

	query := "insert into carts(user_id) values($1) returning *" // create a new cart in case the cart id isnt given

	if err := repo.DB.Raw(query, userId).Scan(&cart).Error; err != nil {
		return models.Cart{}, err
	}
	fmt.Println("cart:", cart)
	return cart, nil

}

// ---------------------------------------- add to cart ----------------------------------------------------------\\

func (repo *cartRepositoryImpl) AddToCart(cart models.CartItems, cartId int) (models.CartItems, error) {
	var body models.CartItems

	query := "insert into cart_items(cart_id,product_id,quantity,amount) values($1,$2,$3,$4) returning *" // insert values into cart_items database

	if err := repo.DB.Raw(query, cartId, cart.ProductId, cart.Quantity, cart.Amount).Scan(&body).Error; err != nil {
		return models.CartItems{}, err
	}
	return body, nil
}

// -------------------------------------------- cart item listing ------------------------------------------------------ \\

func (repo *cartRepositoryImpl) CartItemListing(userId, cartId int) ([]models.CartItems, error) {
	log.Println("******cart item listing*******")
	var carts []models.CartItems

	query := `
	select * from cart_items ct 
	join carts c on c.cart_id = ct.cart_id
	where c.user_id = $1 and c.cart_id = $2
	`

	if err := repo.DB.Raw(query, userId, cartId).Scan(&carts).Error; err != nil {
		return []models.CartItems{}, err
	}
	// fmt.Println("carts", carts)
	return carts, nil
}

// ---------------------------------------------- quantity updation ----------------------------------------------------- \\

func (repo *cartRepositoryImpl) CartItemQuantityUpdations(cartItems, qunatity int) error {

	query := `update cart_items
				set quantity = $1 
				where cart_items_id = $2
	`
	if err := repo.DB.Exec(query, qunatity, cartItems).Error; err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------- cart item deletion ------------------------------------------------------\\

func (repo *cartRepositoryImpl) CartItemDeletion(cartItemsId int) error {

	query := `delete from cart_items where cart_items_id = $1
	`
	if err := repo.DB.Exec(query, cartItemsId).Error; err != nil {
		return err
	}
	return nil
}

// ---------------------------------------- get cart items throught cart items id --------------------------------------- \\

func (repo *cartRepositoryImpl) CartItemsById(cartItemsId int) (models.CartItems, error) {

	var body models.CartItems

	query := `
	select * from cart_items where cart_items_id = $1
	`
	if err := repo.DB.Raw(query, cartItemsId).Scan(&body).Error; err != nil {
		return models.CartItems{}, err
	}
	return body, nil
}

// ---------------------------------------- check if user has carts -------------------------------- \\

func (repo *cartRepositoryImpl) CheckUserCartById(userId int) error {
	var body int

	query := `select count(*) from carts where user_id = $1`

	if err := repo.DB.Raw(query, userId).Scan(&body).Error; err != nil {
		err = errors.New("error in db query to check for carts" + err.Error())
		return err
	}

	if body > 0 {
		return nil
	} else {
		return errors.New("there are no carts for this user")
	}
}

// ------------------------------------------ get carts of user ------------------------------------ \\

func (repo *cartRepositoryImpl) GetCartsByUserId(userInt int) ([]domain.Cart, error) {
	var body []domain.Cart

	query := `
		select * from carts
		where user_id = $1
	`

	if err := repo.DB.Raw(query, userInt).Scan(&body).Error; err != nil {
		err = errors.New("error in db query to get carts" + err.Error())
		return []domain.Cart{}, err
	}
	fmt.Println("body", body)
	return body, nil
}
