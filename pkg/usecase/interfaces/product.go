package interfaces

import "github.com/karthikkalarikal/ecommerce-project/pkg/utils/models"

type ProductUseCase interface {
	ListProducts() ([]models.Product, error)
}
