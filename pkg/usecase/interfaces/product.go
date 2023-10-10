package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type ProductUseCase interface {
	ListProducts() ([]models.Product, error)
	UpdateCategory(category domain.Category, id int) (domain.Category, error)
	DeleteCategory(id int) (domain.Category, error)
	AddProduct(models.Product) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
	AddCategory(domain.Category) (domain.Category, error)
	EditProduct(domain.Product, int) (domain.Product, error)
}
