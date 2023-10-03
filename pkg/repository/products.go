package repository

import (
	"errors"
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	repo *gorm.DB
}

func NewProductRepository(repo *gorm.DB) interfaces.ProductRepository {
	return &productRepositoryImpl{
		repo: repo,
	}
}

func (prod *productRepositoryImpl) ListProducts() ([]models.Product, error) {

	var product_list []models.Product

	query := "SELECT * FROM products"
	err := prod.repo.Raw(query).Scan(&product_list).Error

	if err != nil {
		return []models.Product{}, errors.New("error checking user details")
	}
	fmt.Println(product_list)
	return product_list, nil
}
