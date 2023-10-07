package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type ProductUseCase interface {
	ListProducts() ([]models.Product, error)
	UpdateCategory(category domain.Category) (domain.Category, error)
	DeleteCategory(id int) (domain.Category, error)
}
