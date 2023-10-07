package repository

import (
	"errors"
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
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

// list products
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

// update categories

func (prod *productRepositoryImpl) UpdateCategory(category domain.Category) (domain.Category, error) {

	var body domain.Category

	query := "UPDATE categories SET category_name = $1 WHERE category_id = $2"
	fmt.Println(category.CategoryID, category.CategoryName)

	if err := prod.repo.Exec(query, category.CategoryName, category.CategoryID).Error; err != nil {
		return domain.Category{}, err
	}

	if err := prod.repo.First(&body, category.CategoryID).Error; err != nil {
		return domain.Category{}, err
	}

	return body, nil
}

// delete categories

func (prod *productRepositoryImpl) DeleteCategory(id int) (domain.Category, error) {
	var body domain.Category

	query := "select * from categories where category_id = ?"
	query2 := "delete from categories where id = ?"

	if err := prod.repo.Raw(query2).Scan(&body).Error; err != nil {
		return domain.Category{}, err
	}

	if err := prod.repo.Exec(query, id).Error; err != nil {
		return domain.Category{}, err
	}
	return body, nil
}
