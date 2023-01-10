package router

import (
	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/config"
	"goblog.com/goblog/internal/goblog/handler"
	"goblog.com/goblog/internal/goblog/middleware"
	"goblog.com/goblog/internal/goblog/service"
)

func InitRouter() *gin.Engine {
	return gin.Default()
}

func InitRoutes(
	router *gin.Engine,

	// Handler
	authHandler handler.AuthHandler,
	commentHandler handler.CommentHandler,
	postHandler handler.PostHandler,
	proxyHandler handler.ProxyHandler,

	// Services
	tokenService service.TokenService,

	// Config
	settings *config.Settings,
) any {

	securedApi := router.Group("/api", middleware.JWTAuthMiddleware(tokenService, settings))
	{
		// Auth
		securedApi.POST("/auth/logout", authHandler.Logout)

		// Posts
		postsApi := securedApi.Group("/posts")
		postsApi.GET("", postHandler.GetAll)
		postsApi.GET("/:id", postHandler.GetOne)
		postsApi.POST("", postHandler.Create)
		postsApi.PUT("/:id", postHandler.Update)
		postsApi.PUT("/:id/image", postHandler.AddImage)
		postsApi.DELETE("/:id", postHandler.Delete)

		//Comments
		postsApi.GET("/:id/comments", commentHandler.GetAllForPost)
		postsApi.POST("/:id/comments", commentHandler.Create)
		postsApi.PUT("/:id/comments/:commentId", commentHandler.Update)
	}

	api := router.Group("/api")
	{
		// Auth
		api.POST("/auth/login", authHandler.Login)
		api.POST("/auth/token/refresh", authHandler.RefreshToken)

		// Swagger
		api.Static("/docs", "./assets/swagger-ui")
		api.StaticFile("/docs-src", "./docs/openapi.yml")
	}

	// Minio assets proxy
	router.GET("/assets/images/*proxyPath", proxyHandler.ServePublicBucket)

	return nil
}
