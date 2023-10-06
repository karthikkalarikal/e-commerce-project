package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type AdminUseCase interface {
	UserList() ([]models.UserDetails, error)
	BlockUser(id int, block bool) (domain.Users, error)
}
