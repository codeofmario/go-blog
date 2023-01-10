//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	//_ "github.com/google/wire/cmd/wire"
	"github.com/minio/minio-go/v7"
	"goblog.com/goblog/internal/goblog/config"
	"goblog.com/goblog/internal/goblog/handler"
	"goblog.com/goblog/internal/goblog/repository"
	"goblog.com/goblog/internal/goblog/router"
	"goblog.com/goblog/internal/goblog/service"
	"gorm.io/gorm"
)

func InitRoutes(engine *gin.Engine, settings *config.Settings, db *gorm.DB, store *minio.Client, redis *redis.Client) any {
	wire.Build(
		// Repositories
		repository.NewCommentRepository,
		repository.NewPostRepository,
		repository.NewTokenRepository,
		repository.NewUserRepository,

		// Services
		service.NewAuthService,
		service.NewCommentService,
		service.NewPostService,
		service.NewStoreService,
		service.NewTokenService,
		service.NewUserService,

		// Handlers
		handler.NewAuthHandler,
		handler.NewCommentHandler,
		handler.NewPostHandler,
		handler.NewProxyHandler,

		router.InitRoutes,
	)

	return nil
}
