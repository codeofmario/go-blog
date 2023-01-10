package mapper

import (
	"github.com/google/uuid"
	"goblog.com/goblog/internal/goblog/dto/request"
	"goblog.com/goblog/internal/goblog/dto/response"
	"goblog.com/goblog/internal/goblog/model"
)

func FromDtoToComment(dto *request.CommentRequestDto) *model.Comment {
	id, _ := uuid.Parse(dto.ID)
	postId, _ := uuid.Parse(dto.PostID)

	return &model.Comment{
		Model:  model.Model{ID: id},
		Body:   dto.Body,
		PostID: postId,
	}
}

func FromCommentToDto(model *model.Comment) *response.CommentResponseDto {
	return &response.CommentResponseDto{
		ID:        model.ID.String(),
		Body:      model.Body,
		UserId:    model.UserID.String(),
		PostId:    model.PostID.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
