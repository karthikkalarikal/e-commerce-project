package repository

import (
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
	// fmt.Println("**delete repo")
	query := "delete from users where id = ?"
	// fmt.Println("id:", id)
	err := db.db.Exec(query, id).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// add product
func (db *adminRepositoryImpl) AddProduct(product domain.Product) (domain.Product, error) {
	var products domain.Product

	query := "insert into products (category_id,product_name, product_image,colour,stock,price) values(?,?,?,?,?,?) returning *"

	if err := db.db.Raw(query, product.Category_id, product.ProductName, product.Product_image, product.Colour, product.Stock, product.Price).Scan(&products).Error; err != nil {
		return domain.Product{}, err
	}
	return products, nil
}

// edit product
func (db *adminRepositoryImpl) EditProduct(product domain.Product) (domain.Product, error) {
	var modProduct domain.Product

	query := "UPDATE products SET category_id = ? , product_name = ?, product_image = ?, colour = ?, stock = ?, price = ? WHERE id = ?"

	if err := db.db.Exec(query, product.Category_id, product.ProductName, product.Product_image, product.Colour, product.Stock, product.Price, product.Id).Error; err != nil {
		return domain.Product{}, err
	}

	if err := db.db.First(&modProduct, product.Id).Error; err != nil {
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
