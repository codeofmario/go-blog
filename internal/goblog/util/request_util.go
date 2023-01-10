package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetNamedID(c *gin.Context, name string) uuid.UUID {
	id, _ := uuid.Parse(c.Param(name))
	return id
}

func GetPathID(c *gin.Context) uuid.UUID {
	return GetNamedID(c, "id")
}

func GetUserId(c *gin.Context) uuid.UUID {
	userIdString, _ := c.Get("userId")
	id, _ := uuid.Parse(userIdString.(string))
	return id
}

func GetBody[DTO any](c *gin.Context) *DTO {
	var body *DTO
	c.BindJSON(&body)
	return body
}

func GetBodyAndMapToModel[DTO any, MODEL any](c *gin.Context, mapper func(*DTO) *MODEL) *MODEL {
	body := GetBody[DTO](c)
	return mapper(body)
}
