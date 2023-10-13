package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type UserRepository interface {
	UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error)
	CheckUserAvailability(email string) bool
	UserBlockedStatus(email string) (bool, error)
	FindUserByEmail(email string) (models.UserSignInResponse, error)
	CheckRole(email string) (bool, error)
	AddAddress(models.Address, int) error
	FindAddress(int) ([]models.Address, error)
	SelectAddress(addressId int, val bool) (models.Address, error)
	EditUserDetails(int, models.UserDetailsResponse) (models.UserDetailsResponse, error)
	GetUserDetailsThroughId(int) (models.UserSignInResponse, error)
	ChangeUserPassword(userId int, password string) (models.UserSignInResponse, error)
}
