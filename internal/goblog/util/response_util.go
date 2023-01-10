package util

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"goblog.com/goblog/internal/goblog/dto/response"
	customErrors "goblog.com/goblog/internal/goblog/errors"
)

func CreateResponseIfError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	switch {
	case errors.As(err, &customErrors.BadRequestError{}):
		c.JSON(http.StatusBadRequest, response.ErrorResponseDto{Msg: err.Error()})
	case errors.As(err, &customErrors.UnathorizedError{}):
		c.JSON(http.StatusUnauthorized, response.ErrorResponseDto{Msg: err.Error()})
	case errors.As(err, &customErrors.ForbiddenError{}):
		c.JSON(http.StatusForbidden, response.ErrorResponseDto{Msg: err.Error()})
	case errors.As(err, &customErrors.NotFoundError{}):
		c.JSON(http.StatusNotFound, response.ErrorResponseDto{Msg: err.Error()})
	case errors.As(err, &customErrors.InternalServerError{}):
		c.JSON(http.StatusInternalServerError, response.ErrorResponseDto{Msg: err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, response.ErrorResponseDto{Msg: "Internal server error"})
	}

	c.Abort()
}
