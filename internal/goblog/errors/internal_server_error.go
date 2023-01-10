package errors

type InternalServerError struct {
	Msg string
}

func (e InternalServerError) Error() string {
	return e.Msg
}
