package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type AdminRepository interface {
	UserList() ([]models.UserDetails, error)
}
