package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(User models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
}
