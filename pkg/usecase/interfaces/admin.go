package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminUseCase interface {
	UserList(int, int) ([]models.UserDetailsResponse, error)
	BlockUser(id int) (domain.Users, error)
	FindUser(email string, name string, id string, pageNo int, pageList int) ([]domain.Users, error)
	DeleteUser(id int) (domain.Users, error)
	AddProduct(domain.Product) (domain.Product, error)
	EditProduct(domain.Product) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
	AddCategory(domain.Category) (domain.Category, error)
}
