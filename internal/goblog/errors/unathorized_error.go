package errors

type UnathorizedError struct {
	Msg string
}

func (e UnathorizedError) Error() string {
	return e.Msg
}
