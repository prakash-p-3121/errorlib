package errorlib

import (
	"github.com/gin-gonic/gin"
	restlib "github.com/prakash-p-3121/restlib"
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
	restlib.NotFoundResponse(ctx, err.errorDescription)
}
