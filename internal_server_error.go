package errorlib

import (
	"github.com/gin-gonic/gin"
	restResponder "github.com/prakash-p-3121/rest-response-lib"
)

type InternalServerErrImpl struct {
	AppErrorImpl
}

func NewInternalServerError(desc string) error {
	return &InternalServerErrImpl{AppErrorImpl{
		errorDescription: desc,
	}}
}

func (err *InternalServerErrImpl) Error() string {
	return err.errorDescription
}

func (err *InternalServerErrImpl) SendRestResponse(ctx *gin.Context) {
	restResponder.InternalServerErrorResponse(ctx, err.errorDescription)
}
