package response

type ErrorResponseDto struct {
	Msg    string            `json:"msg"`
	Errors map[string]string `json:"errors,omitempty"`
}
