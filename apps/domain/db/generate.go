package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/hafizdkn/toko-belanja/apps/domain/category"
	"github.com/hafizdkn/toko-belanja/apps/domain/product"
	transactionhistory "github.com/hafizdkn/toko-belanja/apps/domain/transactionHistory"
	"github.com/hafizdkn/toko-belanja/apps/domain/user"
)

func ConnectionDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./toko-belanja.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(user.User{}, category.Category{}, product.Product{}, transactionhistory.Transaction{})
	if err != nil {
		return db, err
	}

	return db, err
}
