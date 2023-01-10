package config

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func InitStore(settings *Settings) *minio.Client {

	client, err := minio.New(settings.MinioAddr, &minio.Options{
		Creds:  credentials.NewStaticV4(settings.MinioAccessKey, settings.MinioSecretKey, ""),
		Secure: false,
	})

	if err != nil {
		panic("Cannot connect to minio")
	}

	return client
}
