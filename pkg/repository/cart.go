package repository

import (
	"fmt"
	"log"

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

// --------------------------------------- cart item listing ------------------------------------------------------\\

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

// ---------------------------------------------- quantity updation -----------------------------------------------------\\
																
func (repo *cartRepositoryImpl) CartItemQuantityUpdations(userId, productId int, qunatity string) error {

	query := `update carts
				set quantity = $1 
				where user_id = $2
				and product_id = $3
	`
	if err := repo.DB.Exec(query, qunatity, userId, productId).Error; err != nil {
		return err
	}
	return nil
}

// ----------------------------------------------- cart item deletion ------------------------------------------------------\\

func (repo *cartRepositoryImpl) CartItemDeletion(userId, productId int) error {

	query := `delete from carts
				where user_id = $1
				and product_id = $2
	`
	if err := repo.DB.Exec(query, userId, productId).Error; err != nil {
		return err
	}
	return nil
}
