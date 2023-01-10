package mapper

import (
	"goblog.com/goblog/internal/goblog/dto/request"
	"goblog.com/goblog/internal/goblog/model"
)

func FromLoginDtoToUser(dto *request.LoginRequestDto) *model.User {
	return &model.User{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
