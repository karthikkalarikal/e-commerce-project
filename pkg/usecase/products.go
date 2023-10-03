package usecase

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	usecase "github.com/karthikkalarikal/ecommerce-project/pkg/usecase/interfaces"
	"github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"
)

type productUseCaseImpl struct {
	usecase interfaces.ProductRepository
}

func NewProductUsecase(usecase interfaces.ProductRepository) usecase.ProductUseCase {
	return &productUseCaseImpl{
		usecase: usecase,
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
