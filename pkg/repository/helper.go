package repository

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type helperRepositoryimpl struct {
	db *gorm.DB
}

func NewHelperRepository(db *gorm.DB) interfaces.HelperRepository {
	return &helperRepositoryimpl{
		db: db,
	}
}

// -------------------------------------helper function that will be useful in different functionalities-----------------------\\
func (repo *helperRepositoryimpl) GetUserDetailsThroughId(id int) (domain.Users, error) {
	var users domain.Users

	query := "select * from users where user_id = $1"

	if err := repo.db.Raw(query, id).Scan(&users).Error; err != nil {
		return domain.Users{}, err
	}
	return users, nil
}

// ----------------------------------------find product by id --------------------------------------------------\\
func (db *helperRepositoryimpl) FindProductById(id int) (domain.Product, error) {
	var product domain.Product

	query := "select * from products where product_id = ?"
	if err := db.db.Raw(query, id).Scan(&product).Error; err != nil {
		return domain.Product{}, err
	}
	fmt.Println(product)
	return product, nil
}

// -------------------------------------- to see if the user id exists  ---------------------------------\\

func (repo *helperRepositoryimpl) FindIfUserExists(userId int, str string) (bool, error) {
	var count int
	query := "select count(*) from "

	query += str // you could give the name of table you want to check
	query += " where user_id = $1"
	if err := repo.db.Raw(query, userId).Scan(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
