package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type OtpRepository interface {
	FindUserByMobileNumber(phone string) bool
	UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error)
}
