package mapper

import (
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/dto/request"
	"goblog.com/goblog/internal/goblog/dto/response"
	"goblog.com/goblog/internal/goblog/model"
)

func FromDtoToPost(dto *request.PostRequestDto) *model.Post {
	return &model.Post{
		Title: dto.Title,
		Body:  dto.Body,
	}
}

func FromPostToDto(model *model.Post) *response.PostResponseDto {
	var imageUrl = ""
	if model.ImageID != uuid.Nil {
		imageUrl = "/assets/images/" + model.ImageID.String()
	}

	return &response.PostResponseDto{
		ID:        model.ID.String(),
		Title:     model.Title,
		Body:      model.Body,
		ImageUrl:  imageUrl,
		UserId:    model.UserID.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
