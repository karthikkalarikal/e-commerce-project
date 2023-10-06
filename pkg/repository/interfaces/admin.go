package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminRepository interface {
	UserList() ([]models.UserDetails, error)
	BlockUser(id int, block bool) (domain.Users, error)
	FindUserByEmail(email string) ([]domain.Users, error)
	DeleteUser(id int) (bool, error)
	AddProduct(domain.Product) (domain.Product, error)
	EditProduct(domain.Product) (domain.Product, error)
	DeleteProduct(id int) (domain.Product, error)
	FindProductById(id int) (domain.Product, error)
}
