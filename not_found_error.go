package errorlib

import (
	"github.com/gin-gonic/gin"
	restResponder "github.com/prakash-p-3121/rest-response-lib"
)

type NotFoundErrorImpl struct {
	AppErrorImpl
}

func NewNotFoundError(desc string) AppError {
	return &NotFoundErrorImpl{AppErrorImpl{
		errorDescription: desc,
	}}
}

func (err *NotFoundErrorImpl) Error() string {

	return err.errorDescription
}

func (err *NotFoundErrorImpl) SendRestResponse(ctx *gin.Context) {
	restResponder.InternalServerErrorResponse(ctx, err.errorDescription)
}
