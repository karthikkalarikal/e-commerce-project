package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockedStatus(email string) (bool, error)
	FindUserByEmail(email string) (models.UserSignInResponse, error)
	CheckRole(email string) (bool, error)
}
