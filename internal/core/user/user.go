package user

import (
	"fmt"
	"golang_auto_shop/internal/core/models"
)

// "gorm.io/driver/postgres"
// "gorm.io/gorm"

type storage interface {
	AddUser(user models.User) error
	DeleteUser(user models.User) error
	UpdateUserByEmail(email string, user models.User) error
	GetUserByEmail(email string) (models.User, error)
	GetUserCarsByEmail(email string) ([]*models.CarModel, error)
}

type UserController struct {
	storage storage
}

func (c *UserController) AddUser(user models.User) error {
	u, err := c.storage.GetUserByEmail(user.Email)
	if err != nil {
		return fmt.Errorf("The request to the repository failed: %v", err)
	}

	if u.Email == user.Email {
		return fmt.Errorf("User with this email already exists")
	}

	err = c.storage.AddUser(user)
	if err != nil {
		return fmt.Errorf("The request to the repository failed: %v", err)
	}

	return nil
} 
