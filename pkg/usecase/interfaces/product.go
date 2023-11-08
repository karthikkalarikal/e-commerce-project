package interfaces

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type ProductUseCase interface {
	ListProducts(int, int) ([]models.Product, error)
	ListProductsByCategory(int) ([]models.Product, error)
	UpdateCategory(category domain.Category, id int) (domain.Category, error)
	DeleteCategory(id int) (domain.Category, error)
	AddProduct(models.Product) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
	AddCategory(domain.Category) (domain.Category, error)
	EditProduct(domain.Product, int) (domain.Product, error)
	AddImage(*gin.Context, *multipart.FileHeader, int) (domain.Image, error)
}
