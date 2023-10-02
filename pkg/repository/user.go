package repository

import (
	"errors"

	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{
		DB: DB,
	}
}

func (u *userDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse

	err := u.DB.Raw("insert into users(name, email, password, phone) values (?,?,?,?) returning id, name, email,phone", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	// fmt.Println(userDetails)
	return userDetails, nil
}

func (c *userDatabase) CheckUserAvailability(email string) bool {
	var count int

	query := "SELECT COUNT(*) from users where email= ?"

	if err := c.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// blocked status
func (c *userDatabase) UserBlockedStatus(email string) (bool, error) {
	var isBlocked bool

	query := "SELECT blocked from users where email = ?"

	if err := c.DB.Raw(query, email).Scan(&isBlocked).Error; err != nil {
		return false, err
	}
	return isBlocked, nil
}

//find user details

func (c *userDatabase) FindUserByEmail(email string) (models.UserSignInResponse, error) {
	var user_details models.UserSignInResponse

	query := "SELECT * FROM users WHERE email = ?"
	err := c.DB.Raw(query, email).Scan(&user_details).Error

	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}
	return user_details, nil
}
