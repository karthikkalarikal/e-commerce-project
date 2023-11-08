package db

import (
	"fmt"

	"github.com/karthikkalarikal/ecommerce-project/pkg/config"
	"github.com/karthikkalarikal/ecommerce-project/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{SkipDefaultTransaction: true})

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Product{})
	db.AutoMigrate(&domain.Category{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.CartItems{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.Payment{})
	db.AutoMigrate(&domain.Coupons{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.Image{})
	return db, err
}
