package repository

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
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

func (repo *cartRepositoryImpl) AddToCart(cart domain.Cart, id int) (domain.Cart, error) {
	var body domain.Cart

	query := "insert into carts(user_id,product_id,quantity,total_price) values($1,$2,$3,$4) returning *"

	if err := repo.DB.Raw(query, id, cart.ProductId, cart.Quantity, cart.TotalPrice).Scan(&body).Error; err != nil {
		return domain.Cart{}, err
	}
	return body, nil
}
