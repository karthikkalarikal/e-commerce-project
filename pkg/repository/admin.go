package repository

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) interfaces.AdminRepository {
	return &adminRepositoryImpl{
		db: db,
	}
}

// view all the users in the database
func (db *adminRepositoryImpl) UserList() ([]models.UserDetails, error) {

	var userList []models.UserDetails

	query := "SELECT * FROM users"

	err := db.db.Raw(query).Scan(&userList).Error

	if err != nil {
		return []models.UserDetails{}, err
	}

	return userList, nil
}

// user block or unblock
func (db *adminRepositoryImpl) BlockUser(id int, block bool) (domain.Users, error) {
	var user domain.Users

	query := "update users set blocked = ? where id = ? returning *"
	err := db.db.Raw(query, block, id).Scan(&user).Error
	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

// search user by email
func (db *adminRepositoryImpl) FindUserByEmail(email string) ([]domain.Users, error) {
	var user []domain.Users

	query := "select * from users where email like ?"

	err := db.db.Raw(query, "%"+email+"%").Scan(&user).Error
	if err != nil {
		return []domain.Users{}, err
	}
	return user, nil
}

// delete user
func (db *adminRepositoryImpl) DeleteUser(id int) (bool, error) {
	fmt.Println("**delete repo")
	query := "delete from users where id = ?"
	fmt.Println("id:", id)
	err := db.db.Exec(query, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
