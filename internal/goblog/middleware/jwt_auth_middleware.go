package middleware

import (
	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/config"
	"goblog.com/goblog/internal/goblog/errors"
	"goblog.com/goblog/internal/goblog/service"
	"goblog.com/goblog/internal/goblog/util"
)

func JWTAuthMiddleware(tokenService service.TokenService, settings *config.Settings) func(c *gin.Context) {
	return func(c *gin.Context) {
		tokenString, err := tokenService.ExtractFromAuthHeader(c.GetHeader("Authorization"))
		if err != nil {
			util.CreateResponseIfError(c, errors.UnathorizedError{Msg: "You are not authorized"})
			return
		}

		claims, err := tokenService.Parse(tokenString, settings.AccessSecret)
		if err != nil {
			util.CreateResponseIfError(c, errors.UnathorizedError{Msg: "You are not authorized"})
			return
		}

		c.Set("userId", claims.Subject)
		c.Next()
	}
}
