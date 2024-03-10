package errorlib

type BadReqError struct {
	AppError
}

func (err *BadReqError) Error() string {
	return err.errorDescription
}

func NewBadReqError(desc string) error {
	return &BadReqError{AppError{
		errorDescription: desc,
	}}
}
