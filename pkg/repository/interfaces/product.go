package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type ProductRepository interface {
	ListProducts() ([]models.Product, error)
}
