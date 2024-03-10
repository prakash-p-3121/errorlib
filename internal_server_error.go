package errorlib

import (
	"github.com/gin-gonic/gin"
	restlib "github.com/prakash-p-3121/restlib"
)

type InternalServerErrImpl struct {
	AppErrorImpl
}

func NewInternalServerError(desc string) AppError {
	return &InternalServerErrImpl{AppErrorImpl{
		errorDescription: desc,
	}}
}

func (err *InternalServerErrImpl) Error() string {
	return err.errorDescription
}

func (err *InternalServerErrImpl) SendRestResponse(ctx *gin.Context) {
	restlib.InternalServerErrorResponse(ctx, err.errorDescription)
}
