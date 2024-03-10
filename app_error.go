package errorlib

/*
AppError instances will not be created at all.
Just the sub-classes of AppError like BadReqError, InternalServerError,
NotFoudError, ConflictError etc.. will be instantiated.
*/
type AppError struct {
	errorDescription string
}

func (err *AppError) Error() string {
	return err.errorDescription
}
