package error

func GeneralRaiseError(text string) error {
	return &GeneralError{text}
}

type GeneralError struct {
	s string
}

func (e *GeneralError) Error() string {
	return e.s
}
