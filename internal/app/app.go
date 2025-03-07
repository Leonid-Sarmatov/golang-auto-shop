package app

import (
	"fmt"
	"golang_auto_shop/internal/storage/gorm"
)

type App struct {
	Storage *gorm.GormAdapter
}

func NewApp() *App {
	return &App{}
}

func (app *App) StartApp() error { 
	storage := gorm.NewGormAdapter()

	err := storage.Init()
	if err != nil {
		return fmt.Errorf("Storage initialization failed: %v", err)
	}

	app.Storage = storage

	return nil
}