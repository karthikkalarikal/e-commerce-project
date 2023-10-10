package repository

import (
	"fmt"
	"strconv"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepositoryImpl struct {
	db   *gorm.DB
	repo interfaces.HelperRepository
}

func NewAdminRepository(db *gorm.DB, repo interfaces.HelperRepository) interfaces.AdminRepository {
	return &adminRepositoryImpl{
		db:   db,
		repo: repo,
	}
}

// ----------------------view all the users in the database------------------------------ \\
func (db *adminRepositoryImpl) UserList(pageList int, offset int) ([]models.UserDetailsResponse, error) {

	var userList []models.UserDetailsResponse

	query := "SELECT * FROM users LIMIT $1 OFFSET $2"

	err := db.db.Raw(query, pageList, offset).Scan(&userList).Error

	if err != nil {
		return []models.UserDetailsResponse{}, err
	}

	return userList, nil
}

// --------------------------------user block or unblock --------------------------------------------------\\
func (db *adminRepositoryImpl) BlockUser(id int, block bool) (domain.Users, error) {
	var user domain.Users

	query := "update users set blocked = ? where user_id = ? returning *"
	err := db.db.Raw(query, block, id).Scan(&user).Error
	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

// -----------------------------------search user--------------------- ---------------------------------------\\
func (db *adminRepositoryImpl) FindUser(email string, name string, id string, pageList int, offset int) ([]domain.Users, error) {
	fmt.Println("***************search repository*******************")
	var users []domain.Users

	args := []interface{}{}
	query := "select * from users where 1=1"

	if email != "" {
		query += " and email like ?"
		searchParam := "%" + email + "%"
		args = append(args, searchParam)
	}
	if name != "" {
		query += " and name like ?"
		searchParam := "%" + name + "%"
		args = append(args, searchParam)
	}

	if id != "" {
		query += " and user_id = ?"
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return []domain.Users{}, err

		}
		searchParam := idInt
		args = append(args, searchParam)
	}

	query += " limit ? offset ?"
	args = append(args, pageList, offset)
	// fmt.Println(query, args)
	err := db.db.Raw(query, args...).Scan(&users).Error
	if err != nil {
		return []domain.Users{}, err
	}

	return users, nil
}

// ----------------------------------delete user -------------------------------------------------\\
func (db *adminRepositoryImpl) DeleteUser(id int) (domain.Users, error) {
	// fmt.Println("**delete repo")
	user, err := db.repo.GetUserDetailsThroughId(id)
	if err != nil {
		return domain.Users{}, err
	}

	query2 := "delete from users where user_id = $1"
	// fmt.Println("id:", id)
	err = db.db.Exec(query2, id).Error
	if err != nil {
		return domain.Users{}, err
	}
	return user, nil
}

// ---------------------check the number of users--------------------- \\
func (db *adminRepositoryImpl) CountUsers() (int, error) {
	var count int

	query := "select count(*) from users"
	if err := db.db.Raw(query).Scan(&count).Error; err != nil {
		return 0, err
	}

	return count, nil

}
