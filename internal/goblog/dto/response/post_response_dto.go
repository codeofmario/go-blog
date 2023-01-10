package response

import "time"

type PostResponseDto struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	ImageUrl  string    `json:"imageUrl"`
	UserId    string    `json:"userId"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
}
