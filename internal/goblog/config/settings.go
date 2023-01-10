package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Settings struct {
	PostgresDsn      string
	MongoDsn         string
	RedisAddr        string
	RedisPassword    string
	MinioAddr        string
	MinioPublicAddr  string
	MinioAccessKey   string
	MinioSecretKey   string
	PublicBucketName string
	PublicKey        string
	PrivateKey       string
	AccessSecret     string
	RefreshSecret    string
}

func InitSettings() *Settings {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Settings{
		PostgresDsn:      os.Getenv("POSTGRES_DSN"),
		RedisAddr:        os.Getenv("REDIS_ADDR"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		MinioAddr:        os.Getenv("MINIO_ADDR"),
		MinioPublicAddr:  os.Getenv("MINIO_PUBLIC_ADDR"),
		MinioAccessKey:   os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecretKey:   os.Getenv("MINIO_SECRET_KEY"),
		PublicBucketName: os.Getenv("PUBLIC_BUCKET_NAME"),
		PublicKey:        os.Getenv("PUBLIC_KEY"),
		PrivateKey:       os.Getenv("PRIVATE_KEY"),
		AccessSecret:     os.Getenv("ACCESS_SECRET"),
		RefreshSecret:    os.Getenv("REFRESH_SECRET"),
	}
}
