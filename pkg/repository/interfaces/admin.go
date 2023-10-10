package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminRepository interface {
	UserList(int, int) ([]models.UserDetailsResponse, error)
	BlockUser(id int, block bool) (domain.Users, error)
	FindUser(email string, name string, id string, pageList int, offset int) ([]domain.Users, error)
	DeleteUser(id int) (domain.Users, error)
	AddProduct(domain.Product) (domain.Product, error)
	EditProduct(domain.Product) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
	FindProductById(id int) (domain.Product, error)
	AddCategory(category domain.Category) (domain.Category, error)
	CountUsers() (int, error)
}
