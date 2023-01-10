package repository

import (
	"fmt"
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/errors"
	"goblog.com/goblog/internal/goblog/model"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAll() ([]*model.Post, error)
	GetOne(id uuid.UUID) (*model.Post, error)
	Create(post *model.Post) (*model.Post, error)
	Update(post *model.Post) (*model.Post, error)
	Delete(id uuid.UUID) (*model.Post, error)
}

type PostRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &PostRepositoryImpl{DB: db}
}

func (r *PostRepositoryImpl) GetAll() ([]*model.Post, error) {
	var posts []*model.Post

	result := r.DB.Find(&posts)
	if result.Error != nil {
		return nil, errors.InternalServerError{Msg: fmt.Sprintf("Not possible to retrive posts")}
	}

	return posts, nil
}

func (r *PostRepositoryImpl) GetOne(id uuid.UUID) (*model.Post, error) {
	var post *model.Post

	result := r.DB.Find(&post, "id = ?", id)
	if result.Error != nil {
		return nil, errors.NotFoundError{Msg: fmt.Sprintf("Post with ID %s was not found", id.String())}
	}

	return post, nil
}

func (r *PostRepositoryImpl) Create(post *model.Post) (*model.Post, error) {

	result := r.DB.Create(&post)
	if result.Error != nil {
		panic(result.Error)
	}

	return post, nil
}

func (r *PostRepositoryImpl) Update(post *model.Post) (*model.Post, error) {

	result := r.DB.Save(&post)
	if result.Error != nil {
		return nil, errors.BadRequestError{Msg: "Cannot create post"}
	}

	return post, nil
}

func (r *PostRepositoryImpl) Delete(id uuid.UUID) (*model.Post, error) {
	post, err := r.GetOne(id)
	if err != nil {
		return nil, err
	}

	result := r.DB.Delete(&post)
	if result.Error != nil {
		return nil, errors.NotFoundError{Msg: fmt.Sprintf("Cannot delete post with ID %s", id.String())}
	}

	return post, nil
}
