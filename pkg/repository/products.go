package repository

import (
	"errors"
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	repo *gorm.DB
}

func NewProductRepository(repo *gorm.DB, helprepo interfaces.HelperRepository) interfaces.ProductRepository {
	return &productRepositoryImpl{
		repo: repo,
	}
}

// ------------------------------------add product -------------------------------------\\

func (db *productRepositoryImpl) AddProduct(product models.Product) (domain.Product, error) {
	var products domain.Product

	query := "insert into products (category_id,product_name, product_image,colour,stock,price) values(?,?,?,?,?,?) returning *"

	if err := db.repo.Raw(query, product.Category_id, product.ProductName, product.Product_image, product.Colour, product.Stock, product.Price).Scan(&products).Error; err != nil {
		return domain.Product{}, err
	}
	return products, nil
}

// -----------------------------------add category ---------------------------------------\\

func (db *productRepositoryImpl) AddCategory(category domain.Category) (domain.Category, error) {

	var adCat domain.Category

	query := "INSERT INTO categories(category_name) VALUES($1) RETURNING * ;"

	if err := db.repo.Raw(query, category.CategoryName).Scan(&adCat).Error; err != nil {
		return domain.Category{}, err
	}

	return adCat, nil
}

// ----------------------------------delete product------------------------------------------\\

func (db *productRepositoryImpl) DeleteProduct(id int) (bool, error) {

	query := "DELETE FROM products WHERE product_id = ?"

	if err := db.repo.Exec(query, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

// --------------------------------------list products --------------------------------------------\\

func (prod *productRepositoryImpl) ListProducts(pageList, offset int) ([]models.Product, error) {

	var product_list []models.Product

	query := "SELECT * FROM products limit $1 offset $2"
	fmt.Println(pageList, offset)
	err := prod.repo.Raw(query, pageList, offset).Scan(&product_list).Error

	if err != nil {
		return []models.Product{}, errors.New("error checking user details")
	}
	fmt.Println("product list", product_list)
	return product_list, nil
}

// --------------------------------------list products by category--------------------------------------\\

func (prod *productRepositoryImpl) ListProductsByCategory(catId int) ([]models.Product, error) {
	var productList []models.Product

	query := `select * from products
			  join categories
		      on products.category_id = categories.category_id 
			  where categories.category_id = $1`

	if err := prod.repo.Raw(query, catId).Scan(&productList).Error; err != nil {
		return []models.Product{}, err
	}

	return productList, nil
}

// update categories

func (prod *productRepositoryImpl) UpdateCategory(category domain.Category, id int) (domain.Category, error) {

	var body domain.Category

	query := "UPDATE categories SET category_name = $1 WHERE category_id = $2"
	// fmt.Println(category.CategoryID, category.CategoryName)

	if err := prod.repo.Exec(query, category.CategoryName, id).Error; err != nil {
		return domain.Category{}, err
	}

	if err := prod.repo.First(&body, id).Error; err != nil {
		return domain.Category{}, err
	}

	return body, nil
}

// ------------------------------------------delete categories -------------------------------------------------\\

func (prod *productRepositoryImpl) DeleteCategory(id int) (domain.Category, error) {
	var body domain.Category

	query := "select * from categories where category_id = ?"
	query2 := "delete from categories where category_id = ?"
	fmt.Println(id)

	if err := prod.repo.Raw(query, id).Scan(&body).Error; err != nil {
		return domain.Category{}, err
	}

	if err := prod.repo.Exec(query2, id).Error; err != nil {
		return domain.Category{}, err
	}
	return body, nil
}

// -------------------------------------------- edit product -----------------------------------------------------\\

func (db *productRepositoryImpl) EditProduct(product domain.Product, id int) (domain.Product, error) {
	var modProduct domain.Product

	query := "UPDATE products SET category_id = ? , product_name = ?, product_image = ?, colour = ?, stock = ?, price = ? WHERE product_id = ?"

	if err := db.repo.Exec(query, product.CategoryId, product.ProductName, product.Colour, product.Stock, product.Price, id).Error; err != nil {
		return domain.Product{}, err
	}

	if err := db.repo.First(&modProduct, id).Error; err != nil {
		return domain.Product{}, err
	}

	return modProduct, nil
}

// ---------------------------------------------- add image ------------------------------------------- \\

func (db *productRepositoryImpl) AddImage(url string, productId int) (domain.Image, error) {
	var body domain.Image

	query := `
	insert into images 
	(product_id,url) 
	values($1,$2)
	returning *
	
	`

	if err := db.repo.Raw(query, productId, url).Scan(&body).Error; err != nil {
		err = errors.New("error in inserting url to image database")
		return domain.Image{}, err
	}

	return body, nil

}

// ----------------------------------------------- display image ----------------------------------- \\

// func (db *productRepositoryImpl) DisplayImage(productId int) (domain.Product, []domain.Image, error) {
// 	var product domain.Product
// 	var images []domain.Image

// 	query1 := `select product_name,product_amount from products
// 	where id  = $1
// 	`
// 	query2 := `select * from images where product_id = $1`

// 	return product, images, nil
// }
