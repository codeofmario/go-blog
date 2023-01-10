package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"goblog.com/goblog/internal/goblog/config"
	"goblog.com/goblog/internal/goblog/errors"
	"mime/multipart"
)

type StoreService interface {
	Save(fileHeader *multipart.FileHeader) (uuid.UUID, error)
	Delete(id uuid.UUID) error
}

type StoreServiceImpl struct {
	store    *minio.Client
	settings *config.Settings
}

func NewStoreService(store *minio.Client, settings *config.Settings) StoreService {
	return &StoreServiceImpl{store: store, settings: settings}
}

func (s *StoreServiceImpl) Save(fileHeader *multipart.FileHeader) (uuid.UUID, error) {
	ctx := context.Background()
	objectName := uuid.New()

	file, err := fileHeader.Open()
	if err != nil {
		return uuid.New(), errors.InternalServerError{Msg: "Cannot store object"}
	}

	_, err = s.store.PutObject(ctx, s.settings.PublicBucketName, objectName.String(), file, fileHeader.Size, minio.PutObjectOptions{ContentType: fileHeader.Header.Get("Content-Type")})
	if err != nil {
		return uuid.New(), errors.InternalServerError{Msg: "Cannot store object"}
	}

	return objectName, nil
}

func (s *StoreServiceImpl) Delete(id uuid.UUID) error {
	err := s.store.RemoveObject(context.Background(), s.settings.PublicBucketName, id.String(), minio.RemoveObjectOptions{GovernanceBypass: true})
	if err != nil {
		return errors.InternalServerError{Msg: "Cannot delete object"}
	}

	return nil
}
