package interfaces

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type OtpRepository interface {
	FindUserByMobileNumber(phone string) (domain.Users, error)
	UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error)
}
