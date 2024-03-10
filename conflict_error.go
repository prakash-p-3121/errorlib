package errorlib

type ConflictError struct {
	AppError
}

func (err *ConflictError) Error() string {
	return err.errorDescription
}

func NewConflictError(desc string) error {
	return &ConflictError{AppError{
		errorDescription: desc,
	}}
}
