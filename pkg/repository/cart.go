package repository

import (
	"fmt"

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

// ----------------------------------------add to cart ----------------------------------------------\\

func (repo *cartRepositoryImpl) AddToCart(cart domain.Cart, userId int, productId int) (domain.Cart, error) {
	var body domain.Cart

	query := "insert into carts(user_id,product_id,quantity,total_price) values($1,$2,$3,$4) returning *"

	if err := repo.DB.Raw(query, userId, productId, cart.Quantity, cart.TotalPrice).Scan(&body).Error; err != nil {
		return domain.Cart{}, err
	}
	return body, nil
}

// --------------------------------------- cart item listing ------------------------------------------------------\\

func (repo *cartRepositoryImpl) CartItemListing(userId int) ([]models.CartItems, error) {
	fmt.Println("******cart item listing*******")
	var carts []models.CartItems

	query := `select ci.product_id,p.product_name,ci.quantity,p.price 
				from carts ci 
				join products p on ci.product_id = p.product_id
				where ci.user_id =  $1
	
	`

	if err := repo.DB.Raw(query, userId).Scan(&carts).Error; err != nil {
		return []models.CartItems{}, err
	}
	// fmt.Println("carts", carts)
	return carts, nil
}

// ----------------------------------------------quantity updation -----------------------------------------------------\\

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
