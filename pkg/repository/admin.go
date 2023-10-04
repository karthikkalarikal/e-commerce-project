package repository

import (
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

func (db *adminRepositoryImpl) UserList() ([]models.UserDetails, error) {

	var userList []models.UserDetails

	query := "SELECT * FROM users"

	err := db.db.Raw(query).Scan(&userList).Error

	if err != nil {
		return []models.UserDetails{}, err
	}

	return userList, nil
}
