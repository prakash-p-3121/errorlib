package errorlib

type NotFoundError struct {
	AppError
}

func (err *NotFoundError) Error() string {

	return err.errorDescription
}

func NewNotFoundError(desc string) error {
	return &NotFoundError{AppError{
		errorDescription: desc,
	}}
}
