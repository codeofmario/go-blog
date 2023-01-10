package handler

import (
	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/dto/request"
	"goblog.com/goblog/internal/goblog/mapper"
	"goblog.com/goblog/internal/goblog/service"
	"goblog.com/goblog/internal/goblog/util"
	"net/http"
)

type AuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthHandlerImpl struct {
	Service service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &AuthHandlerImpl{Service: service}
}

func (h *AuthHandlerImpl) Login(c *gin.Context) {
	user := util.GetBodyAndMapToModel(c, mapper.FromLoginDtoToUser)

	tokens, err := h.Service.Login(user)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (h *AuthHandlerImpl) Logout(c *gin.Context) {
	err := h.Service.Logout(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}

func (h *AuthHandlerImpl) RefreshToken(c *gin.Context) {
	body := util.GetBody[request.TokenRefreshRequestDto](c)
	tokens, err := h.Service.Refresh(body.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusOK, tokens)
}
