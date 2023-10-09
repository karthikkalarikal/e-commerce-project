package usecase

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/karthikkalarikal/ecommerce-project/pkg/helper"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type userUseCaseImpl struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) usecase.UserUseCase {
	return &userUseCaseImpl{
		userRepo: repo,
	}
}

// ------------------------------------------user sign up------------------------------------\\
func (u *userUseCaseImpl) UserSignUp(user models.UserDetails) (models.TokenUsers, error) {
	userExist := u.userRepo.CheckUserAvailability(user.Email)

	fmt.Println("user exists", userExist)
	// fmt.Println("user email", user.Email)
	if userExist {
		return models.TokenUsers{}, errors.New("user already exist, sign in")
	}
	fmt.Println(user)
	if user.Password != user.ConfirmPassword {
		return models.TokenUsers{}, errors.New("password does not match")
	}

	//Hash password since details are validated

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal server error")
	}
	user.Password = string(hashedPassword)

	//add user details to the database
	// fmt.Println("user:", user)

	userData, err := u.userRepo.UserSignUp(user)
	// fmt.Println("userdata", userData)
	if err != nil {
		return models.TokenUsers{}, err
	}

	//jwt token string for user

	tokenString, err := helper.GenerateTokenClients(userData)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error")
	}

	//copies all the details except the password of the user
	var userDetails models.UserDetailsResponse
	err = copier.Copy(&userDetails, &userData)
	// fmt.Println("userDetails:", userDetails)
	// fmt.Println("userData:", userData)

	if err != nil {
		return models.TokenUsers{}, err
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil

}

// ----------------------------------log in------------------------------------------\\
func (u *userUseCaseImpl) LoginHandler(user models.UserLogin) (interface{}, error, bool) {

	ok := u.userRepo.CheckUserAvailability(user.Email)

	if !ok {
		return models.TokenUsers{}, errors.New("the user does not exist"), false
	}
	isBlocked, err := u.userRepo.UserBlockedStatus(user.Email)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal error"), false
	}
	if isBlocked {
		return models.TokenUsers{}, errors.New("user is blocked by admin"), false
	}

	user_details, err := u.userRepo.FindUserByEmail(user.Email)
	// fmt.Println("user details", user_details)
	if err != nil {
		return models.TokenUsers{}, errors.New("internal error"), false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user_details.Password), []byte(user.Password))
	if err != nil {
		return models.TokenUsers{}, errors.New("password incorrect"), false
	}

	if user_details.Role {
		var adminDetails models.AdminDetailsResponse

		adminDetails.Id = user_details.UserID
		adminDetails.Name = user_details.Name
		adminDetails.Email = user_details.Email
		adminDetails.Phone = user_details.Phone
		adminDetails.Role = user_details.Role

		tokenString, err := helper.GenerateTokenAdmin(adminDetails)
		if err != nil {
			return models.TokenUsers{}, errors.New("could not create token due to some internal error"), false
		}

		return models.TokenAdmin{
			Users: adminDetails,
			Token: tokenString,
		}, nil, true
	}

	var userDetails models.UserDetailsResponse

	userDetails.Id = user_details.UserID
	userDetails.Name = user_details.Name
	userDetails.Email = user_details.Email
	userDetails.Phone = user_details.Phone

	tokenString, err := helper.GenerateTokenClients(userDetails)
	if err != nil {
		return models.TokenUsers{}, errors.New("could not create token due to some internal error"), false
	}

	return models.TokenUsers{
		Users: userDetails,
		Token: tokenString,
	}, nil, false

}
