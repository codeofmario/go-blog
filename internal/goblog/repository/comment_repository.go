package repository

import (
	"fmt"

	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/errors"
	"goblog.com/goblog/internal/goblog/model"
	"gorm.io/gorm"
)

type CommentRepository interface {
	GetAllForPost(postId uuid.UUID) ([]*model.Comment, error)
	GetOne(id uuid.UUID) (*model.Comment, error)
	Create(comment *model.Comment) (*model.Comment, error)
	Update(comment *model.Comment) (*model.Comment, error)
}

type CommentRepositoryImpl struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (r *CommentRepositoryImpl) GetAllForPost(postId uuid.UUID) ([]*model.Comment, error) {
	var comments []*model.Comment

	result := r.DB.Where("post_id = ?", postId).Find(&comments)
	if result.Error != nil {
		return nil, errors.InternalServerError{Msg: fmt.Sprintf("Not possible to retrive comments for post with ID %s", postId.String())}
	}

	return comments, nil
}

func (r *CommentRepositoryImpl) GetOne(id uuid.UUID) (*model.Comment, error) {
	var comment *model.Comment

	result := r.DB.Find(&comment, "id = ?", id)
	if result.Error != nil {
		return nil, errors.NotFoundError{Msg: fmt.Sprintf("Comment with ID %s was not found", id.String())}
	}

	return comment, nil
}

func (r *CommentRepositoryImpl) Create(comment *model.Comment) (*model.Comment, error) {

	result := r.DB.Create(&comment)
	if result.Error != nil {
		return nil, errors.BadRequestError{Msg: "Cannot create comment"}
	}

	return comment, nil
}

func (r *CommentRepositoryImpl) Update(comment *model.Comment) (*model.Comment, error) {

	result := r.DB.Save(&comment)
	if result.Error != nil {
		return nil, errors.BadRequestError{Msg: "Cannot update comment"}
	}

	return comment, nil
}
