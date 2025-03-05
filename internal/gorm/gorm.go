package gorm

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter() *GormAdapter {
	return &GormAdapter{}
}

func (adapter *GormAdapter) Init() error {
	dsn := "host=auto-shop-postgres user=auto-shop-admon password=password123 dbname=car_shop port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
		return fmt.Errorf("Can not connect to database: %v", err)
    }
	adapter.db = db
	return nil
}