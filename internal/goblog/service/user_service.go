package service

import (
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/model"
	"goblog.com/goblog/internal/goblog/repository"
)

type UserService interface {
	GetAll() ([]*model.User, error)
	GetOne(id uuid.UUID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(room *model.User) (*model.User, error)
	Delete(id uuid.UUID) (*model.User, error)
}

type UserServiceImpl struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &UserServiceImpl{
		Repo: repo,
	}
}

func (s *UserServiceImpl) GetAll() ([]*model.User, error) {
	return s.Repo.GetAll()
}

func (s *UserServiceImpl) GetOne(id uuid.UUID) (*model.User, error) {
	return s.Repo.GetOne(id)
}

func (s *UserServiceImpl) GetByEmail(email string) (*model.User, error) {
	return s.Repo.GetByEmail(email)
}

func (s *UserServiceImpl) Create(user *model.User) (*model.User, error) {
	return s.Repo.Create(user)
}

func (s *UserServiceImpl) Delete(id uuid.UUID) (*model.User, error) {
	return s.Repo.Delete(id)
}
