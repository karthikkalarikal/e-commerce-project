// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/karthikkalarikal/ecommerce-project/pkg/api"
	"github.com/karthikkalarikal/ecommerce-project/pkg/api/handler"
	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/db"
	"github.com/karthikkalarikal/ecommerce-project/pkg/repository"
	"github.com/karthikkalarikal/ecommerce-project/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	userRepository := repository.NewUserRepository(gormDB)
	helperRepository := repository.NewHelperRepository(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepository, helperRepository)
	userHandler := handler.NewUserHandler(userUseCase)
	otpRepository := repository.NewOtpRepository(gormDB)
	otpUseCase := usecase.NewOtpUsecase(otpRepository, cfg, userRepository)
	otpHandler := handler.NewOtpHandler(otpUseCase)
	productRepository := repository.NewProductRepository(gormDB, helperRepository)
	productUseCase := usecase.NewProductUsecase(productRepository, helperRepository)
	productHandler := handler.NewProductHandler(productUseCase)
	adminRepository := repository.NewAdminRepository(gormDB, helperRepository)
	adminUseCase := usecase.NewAdminUseCase(adminRepository, helperRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase)
	cartRepository := repository.NewCartRepository(gormDB)
	cartUseCase := usecase.NewCartUseCase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)
	orderRepositry := repository.NewOrderRepository(gormDB)
	orderUseCase := usecase.NewOrderUseCase(orderRepositry, helperRepository)
	orderHandler := handler.NewOrderHandler(orderUseCase)
	paymentRepository := repository.NewPaymentRepository(gormDB)
	paymentUseCase := usecase.NewPaymentUseCase(orderRepositry, paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)
	couponRepository := repository.NewCouponRepository(gormDB)
	couponUseCase := usecase.NewCouponUseCase(couponRepository, orderRepositry)
	couponHandler := handler.NewCouponHandler(couponUseCase)
	serverHTTP := http.NewServerHTTP(userHandler, otpHandler, productHandler, adminHandler, cartHandler, orderHandler, paymentHandler, couponHandler)
	return serverHTTP, nil
}
