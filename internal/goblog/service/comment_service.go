package service

import (
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/model"
	"goblog.com/goblog/internal/goblog/repository"
)

type CommentService interface {
	GetAllForPost(postId uuid.UUID) ([]*model.Comment, error)
	GetOne(postId uuid.UUID) (*model.Comment, error)
	Create(comment *model.Comment) (*model.Comment, error)
	Update(comment *model.Comment) (*model.Comment, error)
}

type CommentServiceImpl struct {
	Repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &CommentServiceImpl{Repo: repo}
}

func (s *CommentServiceImpl) GetAllForPost(postId uuid.UUID) ([]*model.Comment, error) {
	return s.Repo.GetAllForPost(postId)
}

func (s *CommentServiceImpl) GetOne(id uuid.UUID) (*model.Comment, error) {
	return s.Repo.GetOne(id)
}

func (s *CommentServiceImpl) Create(comment *model.Comment) (*model.Comment, error) {
	return s.Repo.Create(comment)
}

func (s *CommentServiceImpl) Update(comment *model.Comment) (*model.Comment, error) {
	dbComment, err := s.Repo.GetOne(comment.ID)
	if err != nil {
		return nil, err
	}

	dbComment.Body = comment.Body

	return s.Repo.Update(comment)
}
