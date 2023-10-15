package repository

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type paymentRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRepository(DB *gorm.DB) interfaces.PaymentRepository {
	return &paymentRepositoryImpl{
		DB: DB,
	}
}
