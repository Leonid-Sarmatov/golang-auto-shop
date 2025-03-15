package user

import (
	"fmt"
	"golang_auto_shop/internal/core/models"
)

// "gorm.io/driver/postgres"
// "gorm.io/gorm"

type UserLogicCore interface {
	AddUser(user models.User) error
	DeleteUser(id string) error
	UpdateUserName(id string, name string) error
	UpdateUserEmail(id string, email string) error
	GetUser(id string) (*models.User, error)
	GetUserCars(id string) ([]*models.CarModel, error)
}

type storage interface {
	AddUser(user models.User) error
	DeleteUser(id string) error
	UpdateUserName(id string, name string) error
	UpdateUserEmail(id string, email string) error
	GetUser(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserCars(id string) ([]*models.CarModel, error)
}

type userLogic struct {
	storage storage
}

func NewUserLogicCore() UserLogicCore {
	return &userLogic{}
}

func (c *userLogic) AddUser(user models.User) error {
	if user.Email == "" {
		return fmt.Errorf("Email was empty")
	}

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

func (c *userLogic) DeleteUser(id string) error {
	err := c.storage.DeleteUser(id)
	if err != nil {
		return fmt.Errorf("The request to the repository failed: %v", err)
	}

	return nil
}

func (c *userLogic) UpdateUserName(id string, name string) error {
	if name == "" {
		return fmt.Errorf("Name is empty")
	}

	err := c.storage.UpdateUserName(id, name)
	if err != nil {
		return fmt.Errorf("The request to the repository failed: %v", err)
	}

	return nil
}

func (c *userLogic) UpdateUserEmail(id string, email string) error {
	if email == "" {
		return fmt.Errorf("Name is empty")
	}

	err := c.storage.UpdateUserName(id, email)
	if err != nil {
		return fmt.Errorf("The request to the repository failed: %v", err)
	}

	return nil
}

func (c *userLogic) GetUser(id string) (*models.User, error) {
	u, err := c.storage.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("The request to the repository failed: %v", err)
	}

	return u, nil
}

func (c *userLogic) GetUserCars(id string) ([]*models.CarModel, error) {
	cars, err := c.storage.GetUserCars(id)
	if err != nil {
		return nil, fmt.Errorf("The request to the repository failed: %v", err)
	}

	return cars, nil
}
