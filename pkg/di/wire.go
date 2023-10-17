//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/karthikkalarikal/ecommerce-project/pkg/api"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/db"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewUserRepository,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		http.NewServerHTTP,
		repository.NewOtpRepository,
		usecase.NewOtpUsecase,
		handler.NewOtpHandler,
		repository.NewProductRepository,
		usecase.NewProductUsecase,
		handler.NewProductHandler,
		repository.NewAdminRepository,
		usecase.NewAdminUseCase,
		handler.NewAdminHandler,
		repository.NewCartRepository,
		usecase.NewCartUseCase,
		handler.NewCartHandler,
		repository.NewHelperRepository,
		usecase.NewOrderUseCase,
		handler.NewOrderHandler,
		repository.NewOrderRepository,
		handler.NewPaymentHandler,
		usecase.NewPaymentUseCase,
		repository.NewPaymentRepository,
	)
	return &http.ServerHTTP{}, nil

}
