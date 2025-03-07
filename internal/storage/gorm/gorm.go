package gorm

import (
	"fmt"
	"golang_auto_shop/internal/core/models"

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
	dsn := "host=car-shop-postgres user=car-shop-admin password=password123 dbname=car_shop port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
		return fmt.Errorf("Can not connect to database: %v", err)
    }
	adapter.db = db

	// Автомиграция (создание таблиц, если они не существуют)
    err = db.AutoMigrate(&models.User{}, &models.Engine{}, &models.CarModel{}, &models.UserCar{})
	if err != nil {
		return fmt.Errorf("Automigrate failed: %v", err)
    }
	return nil
}

func (adapter *GormAdapter) Example() {
	// Создаём пользователя
	user := models.User{Name: "John Doe", Email: "john@example.com"}
	adapter.db.Create(&user)

	// Создаём двигатель
	engine := models.Engine{Name: "V8 Turbo", Power: 450}
	adapter.db.Create(&engine)

	// Создаём автомобиль
	car := models.CarModel{Name: "Tesla Model S", Year: 2023, EngineID: engine.ID}
	adapter.db.Create(&car)

	// Связываем пользователя с автомобилем
	adapter.db.Model(&user).Association("Cars").Append(&car)

	// Поиск всех автомобилей пользователя
	var userWithCars models.User
	adapter.db.Preload("Cars").First(&userWithCars, user.ID)
	fmt.Println(userWithCars.Cars) // Выведет список автомобилей пользователя

	// Поиск всех пользователей, владеющих определённым автомобилем
	var carWithUsers models.CarModel
	adapter.db.Preload("Users").First(&carWithUsers, car.ID)
	fmt.Println(carWithUsers.Users) // Выведет список пользователей, владеющих автомобилем
}