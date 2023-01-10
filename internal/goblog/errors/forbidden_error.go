package errors

type ForbiddenError struct {
	Msg string
}

func (e ForbiddenError) Error() string {
	return e.Msg
}
