package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type ProductRepository interface {
	ListProducts(int, int) ([]models.Product, error)
	ListProductsByCategory(int) ([]models.Product, error)
	UpdateCategory(category domain.Category, id int) (domain.Category, error)
	DeleteCategory(id int) (domain.Category, error)
	AddProduct(models.Product) (domain.Product, error)
	DeleteProduct(id int) (bool, error)
	AddCategory(category domain.Category) (domain.Category, error)
	EditProduct(domain.Product, int) (domain.Product, error)
	AddImage(string, int) (domain.Image, error)
}
