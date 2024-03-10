package errorlib

type InternalServerErr struct {
	AppError
}

func (err *InternalServerErr) Error() string {
	return err.errorDescription
}

func NewInternalServerError(desc string) error {
	return &InternalServerErr{AppError{
		errorDescription: desc,
	}}
}
