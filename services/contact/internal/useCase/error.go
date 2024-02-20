package usecase

type Error struct {
	code    int
	message string
}

func newError(message string, code int) *Error {
	return &Error{
		message: message,
		code:    code,
	}
}
