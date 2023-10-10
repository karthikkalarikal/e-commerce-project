package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminUseCase interface {
	UserList(int, int) ([]models.UserDetailsResponse, error)
	BlockUser(id int) (domain.Users, error)
	FindUser(email string, name string, id string, pageNo int, pageList int) ([]domain.Users, error)
	DeleteUser(ctx *gin.Context) (string, error)
	AddProduct(domain.Product) (domain.Product, error)
	EditProduct(domain.Product) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
	AddCategory(domain.Category) (domain.Category, error)
}
