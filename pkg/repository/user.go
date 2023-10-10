package repository

import (
	"errors"
	"fmt"

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

// --------------------User Sign Up----------------------------------------\\
func (u *userDatabase) UserSignUp(user models.UserDetails) (models.UserDetailsResponse, error) {
	var userDetails models.UserDetailsResponse

	err := u.DB.Raw("insert into users(name, email, password, phone) values (?,?,?,?) returning *", user.Name, user.Email, user.Password, user.Phone).Scan(&userDetails).Error

	if err != nil {
		return models.UserDetailsResponse{}, err
	}
	// fmt.Println(userDetails)
	return userDetails, nil
}

// ----------------Check User Availability-----------------------\\
func (c *userDatabase) CheckUserAvailability(email string) bool {
	var count int

	query := "SELECT COUNT(*) from users where email= ?"

	if err := c.DB.Raw(query, email).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// -----------------------------------blocked status--------------------------------------\\
func (c *userDatabase) UserBlockedStatus(email string) (bool, error) {
	var isBlocked bool

	query := "SELECT blocked from users where email = ?"

	if err := c.DB.Raw(query, email).Scan(&isBlocked).Error; err != nil {
		return false, err
	}
	fmt.Println("blocked", isBlocked)
	return isBlocked, nil
}

//----------------------------------------find user details-------------------------------\\

func (c *userDatabase) FindUserByEmail(email string) (models.UserSignInResponse, error) {
	var user_details models.UserSignInResponse

	query := "SELECT * FROM users WHERE email = ?"
	err := c.DB.Raw(query, email).Scan(&user_details).Error
	fmt.Println(user_details)
	if err != nil {
		return models.UserSignInResponse{}, errors.New("error checking user details")
	}
	return user_details, nil
}

// ------------------------------------------find the role--------------------------------------\\

func (c *userDatabase) CheckRole(email string) (bool, error) {
	var isBlocked bool

	query := "SELECT role from users where email = ?"

	if err := c.DB.Raw(query, email).Scan(&isBlocked).Error; err != nil {
		return false, err
	}
	return isBlocked, nil
}

// ------------------------------------------- add address --------------------------------------\\

func (c *userDatabase) AddAddress(address models.Address, userId int) error {

	query := `insert into addresses(user_id,name,house_name,street,city,state,pin) values ($1,$2,$3,$4,$5,$6,$7)`

	if err := c.DB.Exec(query, userId, address.Name, address.HouseName, address.Street, address.City, address.State, address.Pin).Error; err != nil {
		return err
	}

	return nil
}

// ------------------------------------------- get all addresses of user---------------------------\\
func (c *userDatabase) FindAddress(userId int) ([]models.Address, error) {
	var addresses []models.Address

	query := `
				select * from addresses where user_id = $1
	`
	if err := c.DB.Raw(query, userId).Scan(&addresses).Error; err != nil {
		return []models.Address{}, err
	}

	return addresses, nil
}

// ------------------------------------------ select the address --------------------------------------\\
func (c *userDatabase) SelectAddress(addressId int, val bool) (models.Address, error) {
	var address models.Address

	query := `
		update addresses 
		set selection = $1
		where address_id = $2
		returning *
	`
	if err := c.DB.Raw(query, val, addressId).Scan(&address).Error; err != nil {
		return models.Address{}, err
	}
	return address, nil
}
