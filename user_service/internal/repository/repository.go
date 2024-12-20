package repository

import "github.com/faxa0-0/billy/user_service/internal/models"

type UserRepository interface {
	Create(user *models.User) (int, error)
	FindByID(id string) (*models.User, error)
	FindAll() ([]models.User, error)
	Update(id string, user *models.User) error
	Delete(id string) error
}
