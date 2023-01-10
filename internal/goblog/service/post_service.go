package service

import (
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/model"
	"goblog.com/goblog/internal/goblog/repository"
	"mime/multipart"
)

type PostService interface {
	GetAll() ([]*model.Post, error)
	GetOne(id uuid.UUID) (*model.Post, error)
	Create(post *model.Post) (*model.Post, error)
	Update(post *model.Post) (*model.Post, error)
	AddImage(postID uuid.UUID, fileHeader *multipart.FileHeader) (*model.Post, error)
	Delete(id uuid.UUID) (*model.Post, error)
}

type PostServiceImpl struct {
	repo         repository.PostRepository
	storeService StoreService
}

func NewPostService(repo repository.PostRepository, store StoreService) PostService {
	return &PostServiceImpl{repo: repo, storeService: store}
}

func (s *PostServiceImpl) GetAll() ([]*model.Post, error) {
	return s.repo.GetAll()
}

func (s *PostServiceImpl) GetOne(id uuid.UUID) (*model.Post, error) {
	return s.repo.GetOne(id)
}

func (s *PostServiceImpl) Create(post *model.Post) (*model.Post, error) {
	return s.repo.Create(post)
}

func (s *PostServiceImpl) Update(post *model.Post) (*model.Post, error) {
	dbPost, err := s.repo.GetOne(post.ID)
	if err != nil {
		return nil, err
	}

	dbPost.Title = post.Title
	dbPost.Body = post.Body
	return s.repo.Update(dbPost)
}

func (s *PostServiceImpl) AddImage(postID uuid.UUID, fileHeader *multipart.FileHeader) (*model.Post, error) {
	id, err := s.storeService.Save(fileHeader)
	if err != nil {
		return nil, err
	}

	dbPost, err := s.repo.GetOne(postID)
	if err != nil {
		return nil, err
	}

	if dbPost.ImageID != uuid.Nil {
		err := s.storeService.Delete(dbPost.ImageID)
		if err != nil {
			return nil, err
		}
	}

	dbPost.ImageID = id
	return s.repo.Update(dbPost)
}

func (s *PostServiceImpl) Delete(id uuid.UUID) (*model.Post, error) {
	return s.repo.Delete(id)
}
