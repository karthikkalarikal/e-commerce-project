package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type ProductRepository interface {
	ListProducts() ([]models.Product, error)
	UpdateCategory(category domain.Category) (domain.Category, error)
}
