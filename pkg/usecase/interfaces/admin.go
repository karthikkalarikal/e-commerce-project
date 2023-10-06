package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminUseCase interface {
	UserList() ([]models.UserDetails, error)
	BlockUser(id int, block bool) (domain.Users, error)
	FindUserByEmail(ctx *gin.Context) ([]domain.Users, error)
	DeleteUser(ctx *gin.Context) (string, error)
	AddProduct(domain.Product) (domain.Product, error)
}
