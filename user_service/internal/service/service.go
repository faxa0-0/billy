package service

import (
	"github.com/faxa0-0/billy/user_service/internal/models"
	"github.com/faxa0-0/billy/user_service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) (int, error) {
	id, err := s.repo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil

}

func (s *UserService) GetUser(id string) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(id string, user *models.User) error {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
