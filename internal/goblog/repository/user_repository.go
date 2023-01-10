package repository

import (
	"fmt"
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/errors"
	"goblog.com/goblog/internal/goblog/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]*model.User, error)
	GetOne(id uuid.UUID) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(room *model.User) (*model.User, error)
	Delete(id uuid.UUID) (*model.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) GetAll() ([]*model.User, error) {
	var users []*model.User

	result := r.DB.Find(&users)
	if result.Error != nil {
		return nil, errors.InternalServerError{Msg: "Not possible to retrieve users"}
	}

	return users, nil
}

func (r *UserRepositoryImpl) GetOne(id uuid.UUID) (*model.User, error) {
	var user *model.User

	result := r.DB.Find(&user, "id = ?", id.String())
	if result.Error != nil {
		return nil, errors.NotFoundError{Msg: fmt.Sprintf("Not possible to retrive user with id %s ", id)}
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (*model.User, error) {
	var user *model.User

	result := r.DB.Find(&user, "email = ?", email)
	if result.Error != nil {
		return nil, errors.NotFoundError{Msg: fmt.Sprintf("Not possible to retrive user with email %s ", email)}
	}

	return user, nil
}

func (r *UserRepositoryImpl) Create(user *model.User) (*model.User, error) {

	result := r.DB.Create(&user)
	if result.Error != nil {
		return nil, errors.InternalServerError{Msg: "Not possible to create a user"}
	}

	return user, nil
}

func (r *UserRepositoryImpl) Delete(id uuid.UUID) (*model.User, error) {
	user, err := r.GetOne(id)
	if err != nil {
		return nil, err
	}

	result := r.DB.Delete(&user)
	if result.Error != nil {
		return nil, errors.NotFoundError{Msg: fmt.Sprintf("Not possible to delete user with id %s ", user.ID.String())}
	}

	return user, nil
}
