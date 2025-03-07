package models

import (
	"time"
)

// Модель для таблицы users
type User struct {
	ID        uint       `gorm:"primaryKey"`
	Name      string     `gorm:"size:100"`
	Email     string     `gorm:"size:100;unique"`
	Cars      []CarModel `gorm:"many2many:users_cars;"` // Связь многие-ко-многим
	CreatedAt time.Time
}

// Модель для таблицы engines
type Engine struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Power     int
	CreatedAt time.Time
}

type UserCar struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint // Внешний ключ для связи с таблицей users
	CarID     uint // Внешний ключ для связи с таблицей car_models
	CreatedAt time.Time
}

// Модель для таблицы car_models
type CarModel struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Year      int
	EngineID  uint   // Внешний ключ для связи с таблицей engines
	Engine    Engine `gorm:"foreignKey:EngineID"`   // Связь с таблицей engines
	Users     []User `gorm:"many2many:users_cars;"` // Связь многие-ко-многим
	CreatedAt time.Time
}
