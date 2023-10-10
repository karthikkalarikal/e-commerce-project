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

// add product
func (db *adminRepositoryImpl) AddProduct(product domain.Product) (domain.Product, error) {
	var products domain.Product

	query := "insert into products (category_id,product_name, product_image,colour,stock,price) values(?,?,?,?,?,?) returning *"

	if err := db.db.Raw(query, product.CategoryID, product.ProductName, product.Product_image, product.Colour, product.Stock, product.Price).Scan(&products).Error; err != nil {
		return domain.Product{}, err
	}
	return products, nil
}

// edit product
func (db *adminRepositoryImpl) EditProduct(product domain.Product) (domain.Product, error) {
	var modProduct domain.Product

	query := "UPDATE products SET category_id = ? , product_name = ?, product_image = ?, colour = ?, stock = ?, price = ? WHERE id = ?"

	if err := db.db.Exec(query, product.CategoryID, product.ProductName, product.Product_image, product.Colour, product.Stock, product.Price, product.ProductID).Error; err != nil {
		return domain.Product{}, err
	}

	if err := db.db.First(&modProduct, product.ProductID).Error; err != nil {
		return domain.Product{}, err
	}

	return modProduct, nil
}

// delete product
func (db *adminRepositoryImpl) DeleteProduct(id int) (domain.Product, error) {
	var delProduct domain.Product

	delProduct, err := db.FindProductById(id)
	if err != nil {
		return domain.Product{}, err

	}

	query := "DELETE FROM products WHERE id = ?"

	if err = db.db.Exec(query, id).Error; err != nil {
		return domain.Product{}, err
	}
	return delProduct, nil
}

// find product by id
func (db *adminRepositoryImpl) FindProductById(id int) (domain.Product, error) {
	var product domain.Product

	query := "select * from products where id = ?"
	if err := db.db.Raw(query, id).Scan(&product).Error; err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

// add category
func (db *adminRepositoryImpl) AddCategory(category domain.Category) (domain.Category, error) {

	var adCat domain.Category

	query := "INSERT INTO categories(category_name) VALUES($1) RETURNING * ;"

	if err := db.db.Raw(query, category.CategoryName).Scan(&adCat).Error; err != nil {
		return domain.Category{}, err
	}

	return adCat, nil
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
