package response

import "time"

type CommentResponseDto struct {
	ID        string    `json:"id"`
	Body      string    `json:"body"`
	UserId    string    `json:"userId"`
	PostId    string    `json:"postId"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
