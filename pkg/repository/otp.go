package repository

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type otpRepositoryImpl struct {
	DB *gorm.DB
}

func NewOtpRepository(DB *gorm.DB) interfaces.OtpRepository {
	return &otpRepositoryImpl{
		DB: DB,
	}
}

// ----------------------------------find user using moblie number----------------------------\\

func (otp *otpRepositoryImpl) FindUserByMobileNumber(phone string) (domain.Users, error) {
	var count domain.Users
	query := "SELECT * FROM users WHERE phone = ?"
	if err := otp.DB.Raw(query, phone).Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	return count, nil
}

// -------------------------------------user details using phone-----------------------------------\\

func (otp *otpRepositoryImpl) UserDetailsUsingPhone(phone string) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse
	query := "SELECT * FROM users WHERE phone = ?"
	// fmt.Println(phone)
	if err := otp.DB.Raw(query, phone).Scan(&userDetails).Error; err != nil {
		return models.UserDetailsResponse{}, err
	}
	return userDetails, nil
}
