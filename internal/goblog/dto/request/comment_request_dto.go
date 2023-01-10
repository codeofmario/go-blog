package request

type CommentRequestDto struct {
	ID     string `json:"id"`
	Body   string `json:"body"`
	PostID string `json:"postId"`
}
