package usecase

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type productUseCaseImpl struct {
	usecase    interfaces.ProductRepository
	helperRepe interfaces.HelperRepository
}

func NewProductUsecase(usecase interfaces.ProductRepository, helperRepo interfaces.HelperRepository) usecase.ProductUseCase {
	return &productUseCaseImpl{
		usecase:    usecase,
		helperRepe: helperRepo,
	}
}

// ------------------------------add products --------------------------------------\\
func (usecase *productUseCaseImpl) AddProduct(products models.Product) (domain.Product, error) {
	product, err := usecase.usecase.AddProduct(products)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil

}

// -----------------------------add category ----------------------------------------\\
func (usecase *productUseCaseImpl) AddCategory(category domain.Category) (domain.Category, error) {
	adCat, err := usecase.usecase.AddCategory(category)
	if err != nil {
		return adCat, err
	}
	return adCat, nil

}

// ------------------------------delete products ----------------------------------------\\
func (usecase *productUseCaseImpl) DeleteProduct(id int) (domain.Product, error) {
	fmt.Println("******delete repo********")
	var delProduct domain.Product

	delProduct, err := usecase.helperRepe.FindProductById(id)
	if err != nil {
		return domain.Product{}, err

	}

	isTrue, err := usecase.usecase.DeleteProduct(id)
	if err != nil {
		return domain.Product{}, nil
	}
	if isTrue {

		return delProduct, nil
	} else {
		return domain.Product{}, nil
	}
}

func (u *productUseCaseImpl) ListProducts() ([]models.Product, error) {
	productList, err := u.usecase.ListProducts()
	if err != nil {
		return []models.Product{}, err
	}
	// fmt.Println(productList)
	return productList, nil
}

// update category usecase
func (u *productUseCaseImpl) UpdateCategory(category domain.Category, id int) (domain.Category, error) {
	body, err := u.usecase.UpdateCategory(category, id)
	if err != nil {
		return domain.Category{}, err
	}
	return body, nil
}

// delete categories usecase
func (u *productUseCaseImpl) DeleteCategory(id int) (domain.Category, error) {
	body, err := u.usecase.DeleteCategory(id)
	if err != nil {
		return domain.Category{}, err
	}
	return body, nil
}
